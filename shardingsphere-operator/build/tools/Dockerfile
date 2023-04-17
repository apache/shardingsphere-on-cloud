#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

FROM debian:bullseye-slim

ENV MYSQL_VERSION 8.0.32
ENV MYSQL_CLI_DOWNLOAD_URL https://dev.mysql.com/get/Downloads/MySQL-Shell/mysql-shell_8.0.32-1debian11_amd64.deb
ENV ZOOKEEPER_VERSION 3.7.1
ENV ZOOKEEPER_DOWNLOAD_URL https://dlcdn.apache.org/zookeeper/zookeeper-3.7.1/apache-zookeeper-3.7.1-bin.tar.gz
ENV ZOOKEEPER_DIR /app/zookeeper
WORKDIR /app
RUN mkdir -p "/app/start" && chmod -R 777 /app/start
ENTRYPOINT ["sh","-c"]
CMD ["tail -f /dev/null"]
RUN set -eux; \
        \
        apt-get update; \
        apt-get install -y --no-install-recommends ca-certificates wget; \
        apt-get install -y --no-install-recommends curl; \
        apt-get install -y --no-install-recommends libncursesw6 libreadline8 libssh-4; \
        apt-get install -y --no-install-recommends openjdk-11-jre-headless; \
        echo "export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-11-amd64" >> /etc/profile; \
        \
        rm -rf /var/lib/apt/lists/*; \
        \
        wget -O mysql.deb ${MYSQL_CLI_DOWNLOAD_URL}; \
        apt-get install -y ./mysql.deb; \
        rm mysql.deb; \
        wget -O zookeeper.tar.gz "${ZOOKEEPER_DOWNLOAD_URL}"; \
        mkdir -p ${ZOOKEEPER_DIR}; \
        tar -zxf zookeeper.tar.gz -C ${ZOOKEEPER_DIR} --strip-components 1; \
        rm zookeeper.tar.gz; \
        \


