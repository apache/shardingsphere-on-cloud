#
#  Licensed to the Apache Software Foundation (ASF) under one or more
#  contributor license agreements.  See the NOTICE file distributed with
#  this work for additional information regarding copyright ownership.
#  The ASF licenses this file to You under the Apache License, Version 2.0
#  (the "License"); you may not use this file except in compliance with
#  the License.  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.
#

CloudFormation {
  Description "Deploy a ShardingSphere Proxy Cluster in MultiAz mode"

  Parameter("ZookeeperInstanceType") {
    String
    Default "t2.nano"
  }

  Parameter("ShardingSphereInstanceType") {
    String
    Default "t2.micro"
  }

  Parameter("KeyName") {
    String
    Default "test-tf"
    Description "The ssh keypair for remote connetcion"
  }
  
  Parameter("ImageId") {
    Type "AWS::EC2::Image::Id"
    Default "ami-094bbd9e922dc515d"
  }

  Parameter("VpcId") {
    String
    Default "vpc-0ef2b7440d3ade8d5"
    Description "The id of your VPC"
  }

  Parameter("Subnets") {
    Type 'CommaDelimitedList'
    Default "subnet-0f388a6f23063b8c9,subnet-0bc2cd85facb5ca06,subnet-009077567350ef1b7"
    Description "List of subnets sorted by availability zone in your VPC"
  }

  Parameter("SecurityGroupIds") {
    Type 'CommaDelimitedList'
    Default "sg-008e74936b3f9de19"
    Description "List of the id of the SecurityGroups, The security group needs to allow ports 2888, 3888, and 2181 of the zk server to pass through."
  }

  Parameter("HostedZoneName") {
    String
    Default "shardingsphere.org"
    Description "The name of the internal hosted zone, CloudFormation will automatically create `proxy.[InternalHostedZoneName]` for other services to use"
  }

  Parameter("HostedZoneId") {
    String
    Default "Z07043461249YRLI6CRZ8"
    Description "The zone id corresponding to HostedZoneName"
  }

  Parameter("ShardingSpherePort") {
    Integer
    Default 3307
  }

  Parameter("ShardingSphereVersion") {
    String
    Default "5.2.1"
  }

  Parameter("ShardingSphereJavaMemOpts") {
    String
    Default "-Xmx512m -Xms512m -Xmn128m "
  }

  Parameter("ShardingSphereProxyAsgDesiredCapacity") {
    String
    Default "3"
    Description "The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain. see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacitytype, The default value is 3"
  }

  Parameter("ShardingSphereProxyAsgMaxSize") {
    String
    Default "6"
    Description "The maximum size of ShardingSphere Proxy Auto Scaling Group. The default values is 6"
  }

  Parameter("ShardingSphereProxyAsgHealthCheckGracePeriod") {
    Integer
    Default 60
    Description "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check. see https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html"
  }

  Parameter("ZookeeperVersion") {
    String
    Default "3.7.1"
  }

  Parameter("ZookeeperHeap") {
    Integer
    Default 1024
    Description "The maximum heap size given to ZooKeeper"
  }



  (0..2).each do |i| 
    name = "ZK#{i+1}"
    EC2_Instance(name) {
      AvailabilityZone FnSelect(i, FnGetAZs(Ref("AWS::Region")))
      InstanceType Ref("ZookeeperInstanceType")
      ImageId Ref("ImageId")
      KeyName Ref("KeyName")
      SubnetId FnSelect(i, Ref("Subnets"))
      SecurityGroupIds Ref("SecurityGroupIds")
      Tags [ 
        Tag do 
          Key "Name"
          Value "ZK-#{i+1}"
        end
      ]

      (0)

      server = "server.%{idx}=zk-%{idx}.${HostedZoneName}:2888:3888"
      UserData FnBase64(
        FnSub(
          IO.read("./zookeeper-cloud-init.yml"), 
          :SERVERS => FnSub((0..2).map{|i| i == 0 ? server %{:idx => i+1} : ("#{server}" %{:idx => i+1}).insert(0, " " * 4)}.join("\n")), 
          :VERSION => Ref("ZookeeperVersion"),
          :ZK_HEAP => Ref("ZookeeperHeap"),
          :INDEX => i+1,
        )
      )
    }

    domain = "zone#{name}"
    Route53_RecordSet(domain) {
      HostedZoneId Ref("HostedZoneId")
      Name FnSub("zk-#{i+1}.${HostedZoneName}")
      Type "A"
      ResourceRecords [FnGetAtt(name, "PrivateIp")]
      TTL "60"
    }
  end

  launchtemplate_resource_name = "shardingsphereproxyLaunchtemplate"
  launchtemplate_name = "shardingsphereproxy-launchtemplate"

  EC2_LaunchTemplate(launchtemplate_resource_name) {
    LaunchTemplateName launchtemplate_name
    LaunchTemplateData do
      ImageId Ref("ImageId")
      InstanceType Ref("ShardingSphereInstanceType")
      KeyName Ref("KeyName")

      MetadataOptions do
        HttpEndpoint "enabled"
        HttpTokens   "required"
        InstanceMetadataTags "enabled"
      end

      Monitoring do
        Enabled  true
      end

      TagSpecifications [
        {
          :ResourceType => "instance",
          :Tags => [
            {
              :Key => "Name",
              :Value => "shardingsphere-proxy"
            }
          ]
        }
      ]

      UserData FnBase64(
        FnSub(
          IO.read("./shardingsphere-cloud-init.yml"),
          :ZK_SERVERS => FnSub((0..2).map{|i| "zk-#{i+1}.${HostedZoneName}:2181" }.join(",")),
          :VERSION => Ref("ShardingSphereVersion"),
          :JAVA_MEM_OPTS => Ref("ShardingSphereJavaMemOpts")
        )
      )
    end
  }

  ElasticLoadBalancingV2_LoadBalancer("ssinternallb") {
    Name "shardingsphere-internal-lb"
    Scheme "internal"
    Type "network"
    
    mappings = (0..2).map { |x| 
        SubnetMapping do
          SubnetId FnSelect(x, Ref("Subnets"))
        end
    }
    SubnetMappings mappings
    Tags [
      Tag do
        Key "Name"
        Value "shardingsphere"
      end
    ]
  }


  ElasticLoadBalancingV2_TargetGroup("sslbtg") {
    Name "shardingsphere-lb-tg"
    Port Ref("ShardingSpherePort")
    Protocol "TCP"
    VpcId Ref("VpcId")
    TargetGroupAttributes [
      TargetGroupAttribute do
        Key "preserve_client_ip.enabled"
        Value "false"
      end
    ]
    Tags [
      Tag do
        Key "Name"
        Value "shardingsphere"
      end
    ]
  }

  asg_resource_name = "shardingsphereproxyAsg"
  asg_name = "shardingsphereproxy-asg"
  AutoScaling_AutoScalingGroup(asg_resource_name) {
    AutoScalingGroupName asg_name
    AvailabilityZones FnGetAZs(Ref("AWS::Region"))
    DesiredCapacity Ref("ShardingSphereProxyAsgDesiredCapacity")
    MinSize "1"
    MaxSize Ref("ShardingSphereProxyAsgMaxSize")
    HealthCheckGracePeriod  Ref("ShardingSphereProxyAsgHealthCheckGracePeriod")
    HealthCheckType "ELB"

    TargetGroupARNs [ Ref("sslbtg")]

    LaunchTemplate do
      LaunchTemplateName launchtemplate_name
      Version FnGetAtt(launchtemplate_resource_name, "LatestVersionNumber")
    end
  }

  ElasticLoadBalancingV2_Listener("sslblistener") {
    Port Ref("ShardingSpherePort")
    LoadBalancerArn Ref("ssinternallb")
    Protocol "TCP"
    DefaultActions [
      {
        :Type => "forward",
        :TargetGroupArn => Ref("sslbtg")
      }
    ]
  }

  Route53_RecordSet("ssinternaldomain") {
    HostedZoneId Ref("HostedZoneId")
    Name FnSub("proxy.${HostedZoneName}")
    Type "A"
    AliasTarget do 
      HostedZoneId FnGetAtt("ssinternallb", "CanonicalHostedZoneID")
      DNSName FnGetAtt("ssinternallb", "DNSName")
      EvaluateTargetHealth true
    end
  }

  Output("ssinternaldomain") do
    Value Ref("ssinternaldomain")
    Export FnSub("${AWS::StackName}-ShardingSphere-Internal-Domain")
  end

  (0..2).each do |i|
    name = "ZK#{i+1}"
    Output(name) do
      Value FnJoin(":", [Ref(name), FnGetAtt(name, "PrivateIp"), FnGetAtt(name, "AvailabilityZone")])
      Export FnSub("${AWS::StackName}-Zookeeper-Server-#{i+1}")
    end

    zone_name = "zone#{name}"
    Output(zone_name) do
      Value Ref(zone_name)
      Export FnSub("${AWS::StackName}-Zookeeper-Domain-#{i+1}")
    end
  end
}