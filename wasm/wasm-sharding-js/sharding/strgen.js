/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

"use strict";
var is_browser; // true if strgen is being used on a webpage, false if strgen is used via command line
class Strgen {
    constructor() {
        this.pattern = ""; // parameter, the pattern
        this.allow_duplicate_characters = true; // parameter, controls whether the string can be constructed with the same character multiple times (same character at same index) or not
        this.allow_multiple_instances = true; // parameter, allow the string to be constructed with the same character multiple times IF the pattern contains the character more than once
        this.ignore_duplicate_case = false; // parameter, ignore the case of duplicates - i.e. 'A' and 'a' are treated the same, if this is set to true (requires allow_multiple_instances)
        this.allow_logging = false; // parameter, allow the storing of events during the generation process in a list, if set to true
        this.reporting_type = "full"; // parameter, controls level of basic reporting at the start and end of string generation
        this.print_to_console = true; // parameter, allows the log and other output to be printed to the console
        this.error_output_id = "warning"; // parameter, the default UI element where errors will be output (the reference to the element must be the ID)
        this.store_errors = false; // parameter, store errors and warnings in a list of objects when they occur, if set to true
        this.symbol_quantifier_max = 10; // parameter, the highest value possible when using symbol quantifiers
        this.preset = [ // parameter, the character presets and the values of the presets, which can be modified
            {preset_code:"w", value:"_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
            {preset_code:"p", value:"{}[](),./\\:;?!*&@~`'\""},
            {preset_code:"d", value:"0123456789"},
            {preset_code:"c", value:"abcdefghijklmnopqrstuvwxyz"},
            {preset_code:"u", value:"ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
            {preset_code:"l", value:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"},
            {preset_code:"h", value:"0123456789abcdefabcdef"},
            {preset_code:"H", value:"0123456789ABCDEFABCDEF"},
            {preset_code:"o", value:"01234567"},
            {preset_code:"s", value:" "}
        ];
        this.version = "v1.0.0";
    };

    defineVariables() { // defines the non-parameter variables
        this.current_index = -1; // the current pointer/index in the pattern
        this.operators = "[]{}()-\\|/"; // special operator characters responsible for different behaviours TODO: fix error when this is an array
        this.quantifier_operators = [":", ",", "-"]; // operators used within a quantifier, each does the same thing (create range of quantifier values)
        this.symbol_quantifiers = ["+", "*", "?"] // symbol quantifiers based on the quantifiers of regular expression
        this.quantifier_value = 1; // stores the value specified in the pattern within the { }
        this.generated_value_list = []; // where output is stored, to be used in generation at the end of the generation process
        this.temporary_value_list = [];
        this.generated_output = ""; // the full output string
        this.generator_log = []; // if allow_logging is true, events during the generation process will be stored in this list
        this.error_list = []; // if store_errors is true, errors and warnings are stored in this list
        this.error_state = false; // boolean to store whether an error has been encountered
        // assign default values before generation/
        // this fixes a problem with multiple generations with the same instance of the object      
    };

    createString() { // initial method that is called to start generating a random string allow_duplicates = true, allow_logging = false, reporting_type = "full", error_output_id = "warning"
        this.defineVariables();

        if (this.pattern != "") {
            this.checkParameters();
            this.setLogger();
            this.operatorComparison();
            if (this.print_to_console) {
                this.outputLog();
            }
            
            if (!this.error_state) {
                return this.outputString();
            } else {
                this.error_state = false;
                return "";
            }
        } else {
            this.outputError("Pattern is not defined.");
            return "";
        }
    };

    setLogger() { // sets reporting to full if logging is enabled, will post the first log message if reporting is "full" or "less"
        if (this.reporting_type != "full" && this.allow_logging == true) { // set reporting to full if logging is enabled
            this.createLogEntry("Reporting set to full because logging is enabled!");
            this.reporting_type = "full";
        }

        if (this.reporting_type == "full") { // report pattern at the start of generation, if reporting is set to full
            this.createLogEntry("Starting string generation - pattern", this.pattern, true);            
        } else if (this.reporting_type == "less") { // report start of generation, if reporting is set to less
            this.createLogEntry("Strgen-JS - start", undefined, true);
        }
    };

    checkParameters() {
        if (this.allow_multiple_instances == true && this.ignore_duplicate_case == true) {
            this.outputWarning("Cannot ignore character case of duplicates, if multiple instances of characters are allowed!");
            this.ignore_duplicate_case = false;
        }
    };

    lookahead() { // return the next character in the string
        return this.pattern.charAt(this.current_index + 1);
    };

    next() { // increment the current_index value and return the character at that value in the string
        this.current_index += 1;
        return this.pattern.charAt(this.current_index);
    };

    current() { // return the character at the current_index position in the string
        return this.pattern.charAt(this.current_index);
    };

    last() { // return the character before the current_index position in the string
        return this.pattern.charAt(this.current_index - 1);
    };

    operatorComparison() { // main method which is used to determine whether the current character is an operator or not 
                           // (the first character is usually an operator if trying to generate random strings)
        this.next();
        this.createLogEntry("Parsing character at position " + (this.current_index + 1));

        if (this.operators.includes(this.current()) == true || this.symbol_quantifiers.includes(this.current()) == true) {
            this.determineOperator(this.pattern.charAt(this.current_index));
        } else {
            if (!this.symbol_quantifiers.includes(this.lookahead())) { // catch and process literal and call operatorComparison() again
                this.getLiteral();
            } else {
                this.addCharToList(this.current());
            }

            this.operatorComparison();
        }
    };

    determineOperator(operator) { // if an operator was found in operatorComparison, find out what operator it is and respond
        if (operator != "") {
            this.createLogEntry("Operator", operator);
            switch(operator) {
                case "[":
                    this.getCharacterSet();
                    break;
                case "]":
                    this.createLogEntry("End of range reached", this.generated_value_list.toString());
                    if (this.lookahead() != '{' && !this.symbol_quantifiers.includes(this.lookahead())){
                        this.buildGeneratedString(this.selectValueFromList(1, undefined, this.allow_duplicate_characters));    
                    }
                    break;
                case "{":
                    this.quantifier_value = this.getQuantifier();
                    break;
                case "}":
                    if (this.quantifier_value == 0) {
                        this.createLogEntry("End of quantifier reached", "0");
                        this.createLogEntry("No generation as quantifier is 0");
                        this.createLogEntry("Clearing values list...");
                        this.generated_value_list = [];  
                    } else {
                        this.createLogEntry("End of quantifier reached", this.quantifier_value);
                        this.createLogEntry("Contents of value list", this.generated_value_list.toString());
                        this.buildGeneratedString(this.selectValueFromList(this.quantifier_value, undefined, this.allow_duplicate_characters));
                    }
                    this.quantifier_value = 1;
                    break;
                case '(':
                    this.generateSequence();
                    break;
                case ')':
                    this.createLogEntry("End of sequence reached");

                    if (this.temporary_value_list.length >= 1) {
                            this.generated_value_list = this.generated_value_list.concat(this.temporary_value_list);
                            this.temporary_value_list = [];
                    }

                    if (this.lookahead() != '{' && !this.symbol_quantifiers.includes(this.lookahead())){
                        this.createLogEntry("Contents of value list", this.generated_value_list.toString());
                        this.buildGeneratedString(this.selectValueFromList(1, undefined, false));  
                    }
                    break;
                case '/':
                    this.next();
                    this.getLiteral();
                    break;
                case '+':
                case '*':
                case '?':
                    this.quantifier_value = this.getQuantifier();

                    if (this.quantifier_value == 0) {
                        this.createLogEntry("Symbol quantifier reached", "0");
                        this.createLogEntry("No generation as quantifier is 0");
                        this.createLogEntry("Clearing values list...");
                        this.generated_value_list = [];  
                    } else {
                        this.createLogEntry("Symbol quantifier reached", this.quantifier_value);
                        this.createLogEntry("Final contents of value list", this.generated_value_list.toString());
                        this.buildGeneratedString(this.selectValueFromList(this.quantifier_value, undefined, this.allow_duplicate_characters));
                    }
                    this.quantifier_value = 1;
                    break;
                default:
                    if (!this.symbol_quantifiers.includes(this.lookahead())) {
                        this.getLiteral();
                    } else {
                        this.addCharToList(this.current());
                    }
                    break;
            }
            this.operatorComparison();
        } else {
            if (this.reporting_type == 'full') { // report the full information at the end, if reporting is set to full
                if (this.error_state) {
                    this.createLogEntry("End of pattern reached - string failed to generate due to error", undefined, true);
                } else if (this.generated_output == "") {
                    this.createLogEntry("End of pattern reached - no string generated", undefined, true);
                } else {
                    this.createLogEntry("End of pattern reached - final generated string", this.outputString(), true);  
                } 
            } else if (this.reporting_type == "less") { // report complete generation at the end, if reporting is set to less
                this.createLogEntry("Strgen-JS - finish", undefined, true);
            }
        }
    };

    getCharacterSet() { // if the operator was the start of a character class definition, begin reading the pattern, and react according to a set of comparisons
        do { // starts at the [ which was successfully read in the stage before this function
            this.createLogEntry("Processing range at pattern position " + (this.current_index + 1));

            var current_character = this.next();

            if (this.current() == '\\') {
                this.createLogEntry("Preset character at position " + (this.current_index + 1) + ", getting values for preset", this.lookahead());
                this.getPresetValues(this.next());
                this.generateRangeValue();
            } else if (this.operators.includes(this.current()) == true && this.last() != '/' && this.current() != '/') {// if the current character is an unbroken operator, throw error
                if(this.pattern.charAt(this.current_index) != "") {
                    this.outputError("Unexpected operator at position " + (this.current_index + 1) + ", operator '" + this.pattern.charAt(this.current_index) + "'.");
                } else {
                    this.outputError("Character class not closed.");
                }
                break;
            } else if (this.lookahead() == '-' && current_character != '/') { // if the next character is an unbroken '-' operator
                this.createLogEntry("Unbroken operator", "-");

                var character_store;
                character_store = current_character; // take current character (left side of hyphen) and store it temp
                this.next(); // skip hyphen

                if (this.lookahead() == '/') { // if character after hyphen is / break character
                    this.next() // skip \
                    current_character = this.next(); // character after the /
                } else {
                    current_character = this.next(); // assign the next character (right side of hyphen) and store it as the current character
                }
                this.createLogEntry("Range found", character_store + " , " + current_character);

                this.generateRangeValue(character_store.charCodeAt(0), current_character.charCodeAt(0)); // generate_range_value(ascii value of left side , ascii value of right side)
            } else if (this.lookahead() != '-' && current_character != '/') { // if the next character is not the "-" operator, and the current isn't a character break, then push current()
                this.createLogEntry("Literal added to range", current_character);
                this.generated_value_list.push(this.current());
            } 
        } while (this.lookahead() != ']')
    };

    getLiteral() { // output a literal character, skipping any generation
        this.createLogEntry("Literal", this.pattern.charAt(this.current_index));
        this.buildGeneratedString(this.pattern.charAt(this.current_index));
    };

    addCharToList(char) {
        this.createLogEntry("Adding literal '" + char + "' to values list");
        this.generated_value_list.push(char);
    };

    getQuantifier() { // get the value within quantifier operators, if present
        this.createLogEntry("Processing quantifier at pattern position " + (this.current_index + 1));
        var start_value = this.current_index + 1;
        var quantifier_value;
        var quantifier_first_value;
        var quant_range_state = false;
        var symbol_quantifier = false;

        if (this.symbol_quantifiers.includes(this.current())) {
            this.createLogEntry("Symbol quantifier specified");
            quant_range_state = true;
            symbol_quantifier = true;
            if(this.current() == "?") {
                this.createLogEntry("Quantifier", "?");
                quantifier_value = 1;
                quantifier_first_value = 0;
            } else if (this.current() == "*") {
                this.createLogEntry("Quantifier", "*");
                quantifier_value = this.symbol_quantifier_max;
                quantifier_first_value = 0;
            } else if (this.current() == "+") {
                this.createLogEntry("Quantifier", "+");
                quantifier_value = this.symbol_quantifier_max;
                quantifier_first_value = 1;
            }
        } else {
            do {
                if (this.operators.includes(this.lookahead()) == false && this.quantifier_operators.includes(this.lookahead()) == false) { // if lookahead is not any operator
                    if (quantifier_value == undefined) {
                        quantifier_value = this.next();
                    } else {
                        quantifier_value+= this.next();  
                    }
                } else if (this.quantifier_operators.includes(this.lookahead()) == true && quantifier_first_value == undefined) { // if lookahead is quantifier operator i.e. , : -
                    quant_range_state = true;
                    this.createLogEntry("Quantifier range specified");
                    quantifier_first_value = quantifier_value;
                    quantifier_value = "";
                    this.next();
                } else if (this.lookahead() == "") { // if lookahead is nothing (error case)
                    this.outputError("Quantifier not closed.");
                    break;  
                } else {
                    this.next();
                    quantifier_value = 1;
                    this.outputError("Unexpected character at position " + (this.current_index + 1) + ", character '" + this.pattern.charAt(this.current_index) + "'.");
                    break;          
                }
            } while (this.lookahead() != '}')           
        }
        
        if (quant_range_state == true) {

            if (quantifier_first_value == undefined || quantifier_first_value == "") {
                quantifier_first_value = 0;
            } else if (quantifier_value == undefined || quantifier_value == "") {
                this.outputWarning("Max quantifier value was not set, quantifier at position " + start_value + " set to 0.");
                quantifier_value = 0;
            }

            this.createLogEntry("Generating random quantifier between", quantifier_first_value + " and " + quantifier_value);

            quantifier_first_value = parseInt(quantifier_first_value);
            quantifier_value = parseInt(quantifier_value);

            if (quantifier_first_value > quantifier_value) { // swap values if the quantifier_first_value is the largest value
                var store = quantifier_first_value;
                quantifier_first_value = quantifier_value;
                quantifier_value = store;
                this.createLogEntry("Quantifier values swapped");
            }

            // temporary solution to getting a random quantifier from a range
            var quantifier_array = [];

            for (var count = quantifier_first_value; count <= quantifier_value; count++) { // populate array with every value between quantifier_first_value and quantifier_value
                quantifier_array.push(count);
            }
            this.createLogEntry("Quantifier range values", quantifier_array.toString());

            var random_value = Math.floor(Math.random() * quantifier_array.length);
            var selected_quantifier = quantifier_array[random_value]; // select a value based math.random and array length
            this.createLogEntry("Selected index", random_value + ", selected quantifier value: " + selected_quantifier);

            quantifier_array = [];
            //end of temp code 

            quantifier_value = selected_quantifier;
        }

        if (this.allow_duplicate_characters == false) {
            var value_list_length = this.getValueListLength();
            if (quantifier_value > value_list_length) {
                this.outputWarning("Character quantifier at position " + start_value + " reduced from " + 
                    quantifier_value + " to " + value_list_length + 
                    ". Toggle 'Allow Duplicate Characters' to generate the full amount.")

                    quantifier_value = value_list_length;
            }
        }

        this.createLogEntry("Quantifier value is " + quantifier_value);

        if (quantifier_value == 0 && quant_range_state == false) {
            this.createLogEntry("No value was returned. Character quantifier at position " + start_value + " is 0.", undefined, true);
        } else if (quantifier_value == 0 && quant_range_state == true) {
            this.createLogEntry("No value was returned. Character quantifier range at position " + start_value + " generated the value 0!", undefined, true);
        }

        if (isNaN(quantifier_value)) {
            this.outputError("Quantifier at position " + start_value + " contains invalid characters.");
        }

        return parseInt(quantifier_value);
    };

    generateRangeValue(first_value, second_value, character_index = first_value) { // generate all possible values in the user defined range (defined with '-' character)
        if (first_value > second_value) { // swap values if the firstvalue is the largest value
            var store = first_value;
            first_value = second_value;
            second_value = store;
            character_index = first_value;
            this.createLogEntry("Range values swapped");
        }

        if (character_index <= second_value && character_index >= first_value) { // if character_index is within the range specified
            this.generated_value_list.push(String.fromCharCode(character_index));
            this.generateRangeValue(first_value, second_value, parseInt(character_index+=1));
        }
    };

    getPresetValues(character, count = 0) { // determine what set of pre-defined values will be used when generating with the '\' symbol, then call generatePresetValues
        if (count == this.preset.length) {
            this.getLiteral(character);
            this.outputWarning("Invalid preset range. \'\\" + character + "\' is not a valid preset.");
        } else if (this.preset[count].preset_code == character) {
            this.createLogEntry("Found preset", JSON.stringify(this.preset[count]));
            this.generatePresetValues(this.preset[count].value);
        } else {
            count = count + 1;
            this.getPresetValues(character, count);
        }
    };

    generatePresetValues(preset_values, character_index = 0) { // split the preset_characters string and push each individual character into the values array
        if (preset_values != undefined && character_index < preset_values.length) {
            this.generated_value_list.push(preset_values.charAt(character_index));
            this.generatePresetValues(preset_values, character_index+=1);
        }
    };

    generateSequence() { // split each value in the sequence and push it to the generated_value_list array
        var string_value = "";
        var last_operator = "none";
        var temp_string;
        this.createLogEntry("Processing sequence at pattern position " + (this.current_index + 1));

        while(this.current() != ')') { // while the current character is not the end of the sequence
            if (this.lookahead() == '|' || this.lookahead() == ')' || this.lookahead() == '&') { // if the next character is a closing bracket or sequence operators
                if (this.lookahead() == '|' && last_operator != "&" || last_operator == '|' && this.lookahead() == ')') {
                    // if next character is OR operator and last_operator is not AND, or, if last_operator is OR operator and next character is end of sequence - perform OR
                    this.createLogEntry("OR operator - last operator", last_operator);
                    last_operator = '|';

                    if (string_value != "") {
                        this.temporary_value_list.push(string_value);
                        this.createLogEntry("OR word parsed", string_value);
                        string_value = "";                       
                    }

                    if (this.lookahead() == ')') { break; } else { this.next(); }
                } else if (this.lookahead() == '&' || last_operator == '&' && this.lookahead() == ')' || last_operator == '&' && this.lookahead() == '|') {
                    // if next character is AND operator, or, if last_operator is AND operator and next character is end of sequence or OR operator - perform AND
                    this.createLogEntry("AND operator - last operator", last_operator);
                    last_operator = "&";

                    if (temp_string == undefined) {
                        temp_string = string_value;
                    } else {
                        temp_string += string_value;
                    }

                    this.createLogEntry("AND word parsed", string_value);
                    string_value = "";

                    if (this.lookahead() == ')' || this.lookahead() == '|') { // if next character is end or OR operator (AND operator has ended/no more ANDs yet) - split temp_string and create random word
                        this.createLogEntry("OR operator or end ahead");
                        var output_string = "";
                        var temp_string_length = temp_string.length;
                        var temp_string_array = temp_string.split("");

                        this.createLogEntry("Combined string character range", temp_string_array.toString());

                        this.generated_value_list.push(this.generateAndString(temp_string_array));
                        temp_string = "";

                        if (this.lookahead() == ')') { // if next character is end of sequence, break loop
                            break;  
                        } else if (this.lookahead() == '|') { // if next character is OR operator, set last_operator to OR and move to next character
                            if (last_operator = '&') {
                                this.createLogEntry("Storing " + this.generated_value_list.toString() + " in a different list temporarily");
                                this.temporary_value_list = this.temporary_value_list.concat(this.generated_value_list);
                                this.generated_value_list = [];
                            }
                            last_operator = '|'
                            this.createLogEntry("last_operator set to '|'");
                            this.next();
                        }
                    } else { 
                        this.next(); 
                    }
                }
                else if (this.lookahead() == ')' && last_operator == "none" || this.lookahead() == '') {
                    if (string_value == "") {
                        this.outputWarning("Unbroken Sequence starting at position " + (this.current_index + 1) + " does not contain any values.");
                    } else {
                        this.generated_value_list.push(string_value);
                        //this.outputWarning("Sequence starting at position " + ((this.current_index + 1) - string_value.length) + " only contains one value.")         
                    }
                    break;
                }
            }
            else if (this.lookahead() != '') {
                this.next();
                if (this.operators.includes(this.current()) == true || this.symbol_quantifiers.includes(this.current()) == true) {
                    if (this.current() == "[") {
                        this.getCharacterSet();
                    } else if (this.current() == "]") {
                        this.createLogEntry("End of range reached", this.generated_value_list.toString());
                        if (this.lookahead() != '{' && !this.symbol_quantifiers.includes(this.lookahead())) {
                            string_value += this.selectValueFromList(this.quantifier_value, undefined, this.allow_duplicate_characters);
                        }  
                    } else if (this.current() == "{") {
                        this.quantifier_value = this.getQuantifier();
                    } else if (this.current() == "}") {
                        if (this.quantifier_value == 0) {
                            this.createLogEntry("End of quantifier reached", "0");
                            this.createLogEntry("No generation as quantifier is 0");  
                        } else {
                            this.createLogEntry("End of quantifier reached", this.quantifier_value);
                            this.createLogEntry("Contents of value list", this.generated_value_list.toString());

                            string_value += this.selectValueFromList(this.quantifier_value, undefined, this.allow_duplicate_characters);
                        }
                        this.quantifier_value = 1;
                        this.generated_value_list = [];
                    } else if (this.symbol_quantifiers.includes(this.current())) {
                        this.quantifier_value = this.getQuantifier(this.current());

                        string_value += this.selectValueFromList(this.quantifier_value, undefined, this.allow_duplicate_characters);

                        this.quantifier_value = 1;
                        this.generated_value_list = [];
                    } else if (this.current() == "/") {
                        string_value += this.next();
                    }
                } else {
                    string_value += this.current();                   
                }
            } else {
                this.outputError("End of sequence expected at position " + (this.current_index + 1) + ".");
                break;
            }
        }
    };

    generateAndString(values_array, index = 0, output_string = "", array_original_length) {
        if (array_original_length == undefined) {
            array_original_length = values_array.length;
        }

        if (array_original_length > index) {
            var random_value = Math.floor(Math.random() * values_array.length);

            this.createLogEntry("Sequence processing action " + (index + 1) + " - Selected sequence character", values_array[random_value]);

            output_string += values_array[random_value];
            values_array.splice(random_value, 1);

            this.createLogEntry("Sequence processing action " + (index + 1) + " - Range after selection", values_array.toString());
            return this.generateAndString(values_array, index += 1, output_string, array_original_length);
        } else {
            this.createLogEntry("AND string generated", output_string);
            return output_string;
        }
    };

    selectValueFromList(defined_no_of_chars, character_index = 0, allow_duplicates = false, output = "") { // pick a random value from the generated_value_list and do this quantifier_value number of times
        if (character_index < this.quantifier_value) {
            var random_value = Math.floor(Math.random() * this.generated_value_list.length);

            if (this.generated_value_list[random_value] != undefined) {
                output += this.generated_value_list[random_value]
                this.createLogEntry("Selected value", this.generated_value_list[random_value]);
            } else {
                if (output == "") {
                    this.outputWarning("No value was returned. Please check the template.");
                    if (this.generated_value_list.length == 0) {
                        this.createLogEntry("<b>Ending value selection and continuing generation...</b>");
                        character_index = this.quantifier_value;
                    }                 
                }
            }

            if (allow_duplicates == false) {
                var value = this.generated_value_list[random_value];

                this.removeValueFromList(value, random_value, true);
                if (this.allow_multiple_instances == false) {
                    this.removeValueFromList(value);
                    if (this.ignore_duplicate_case == true && value.match(/[a-zA-Z]/)) {
                        var value_upper = value.toUpperCase();
                        var value_lower = value.toLowerCase();

                        if (value != value_lower) {
                            value = value_lower;
                        } else {
                            value = value_upper;
                        }

                        this.removeValueFromList(value);
                    }
                }
            }

            return this.selectValueFromList(this.quantifier_value, character_index+=1, allow_duplicates, output);
        } else {
            this.generated_value_list = [];
            return output;
        }
    };

    removeValueFromList(value, index = 0, one_value_only = false, previously_searched = false, count = 0) {
        if (this.generated_value_list.indexOf(value) != -1) {
            var value_index = this.generated_value_list.indexOf(value, index);

            this.generated_value_list.splice(value_index, 1);
            count += 1;

            if (one_value_only == false) {
                this.removeValueFromList(value, 0, false, true, count);
            } else {
                this.createLogEntry("Removed value '" + value + "' from array", this.generated_value_list.toString());
            }
        } else if (this.generated_value_list.indexOf(value) == -1 && previously_searched == true) {
            if (count > 1) {
                this.createLogEntry("Removed value '" + value + "' from " + count + " indexes in the array", this.generated_value_list.toString());
            } else {
                this.createLogEntry("Removed value '" + value + "' from array", this.generated_value_list.toString());
            }  
        } else {
            this.createLogEntry("Value '" + value + "' was not found in the values list", this.generated_value_list.toString());
        }
    };

    getValueListLength(index = 0, count_list = [], current_count = 0, allow_multiple = this.allow_multiple_instances, ignore_case = this.ignore_duplicate_case) {
        if (allow_multiple == true && ignore_case == false) {
            this.createLogEntry("Values array length is", this.generated_value_list.length);
            return this.generated_value_list.length;
        } else if (index != this.generated_value_list.length && ignore_case == true) {
            if (count_list.indexOf(this.generated_value_list[index].toLowerCase()) == -1) {
                if (count_list.indexOf(this.generated_value_list[index].toUpperCase()) == -1) {
                    count_list.push(this.generated_value_list[index]);
                }
            }
            index += 1;
            return this.getValueListLength(index, count_list, count_list.length);
        } else if (index != this.generated_value_list.length) {
            if (count_list.indexOf(this.generated_value_list[index]) == -1) {
                count_list.push(this.generated_value_list[index]);
            }
            index += 1;
            return this.getValueListLength(index, count_list, count_list.length);
        } else if (index == this.generated_value_list.length) {
            this.createLogEntry("List counted, " + current_count  + " unique values. Unique values are", count_list.toString());
            return current_count;
        }
    };

    buildGeneratedString(output) { // construct the generated string as every value is selected, and store it in generated_output
        this.generated_output += output;
    };

    outputString() { // output the completed string at the end of execution
        return this.generated_output;
    };

    outputWarning(message) { // output a warning to the UI element, to the console, or store it in error_list
        if (is_browser == true && document.getElementById(this.error_output_id)) {
            document.getElementById(this.error_output_id).innerHTML += message;
        } else if (this.store_errors == true) {
            this.error_list.push({
                msg: message,
                state: "warning"
            });
        } else {
            console.error(message);
        }

        this.createLogEntry("<b>WARNING</b>", message, true);
    };

    outputError(message) { // output an error to the console and set error_state to true, and either display the error message on the UI element, or store it in error_list
        if (is_browser == true && document.getElementById(this.error_output_id)) {
            document.getElementById(this.error_output_id).innerHTML += message;
        } else if (this.store_errors == true) {
            this.error_list.push({
                msg: message, 
                state: "error"
            });
        } 

        console.error(message);
        this.createLogEntry("<b>ERROR</b>", message, true);

        this.error_state = true;    
    };

    createLogEntry(caption, content = undefined, enabled = this.allow_logging) { // create a new log entry
        if(enabled == true && this.reporting_type != "none") {
            var timestamp = new Date();
            var timestamp_text = timestamp.toTimeString().split(" ")[0] + ":" + timestamp.getMilliseconds();
            var log_entry;

            if(content != undefined) {
                if(content != "") {
                    if(content != " ") {
                        log_entry = timestamp_text + " - " + caption + ": " + content;
                    } else {
                        log_entry = timestamp_text + " - " + caption + ": white space";
                    }
                } else {
                    log_entry = timestamp_text + " - " + caption + ": empty";
                }
            } else {
                log_entry = timestamp_text + " - " + caption;
            }

            this.generator_log.push(log_entry);
        }
    };

    outputLog() { // output the log to the console
        for (var count = 0; count <= this.generator_log.length - 1; count++) {
            console.log((count + 1) + " - " + this.generator_log[count]);
        }
    };

    printVersion() { // print strgen version
        console.log("strgen-js " + this.version);
    };
};

if (typeof module !== 'undefined' && typeof module.exports !== 'undefined') {
    module.exports = Strgen;
    is_browser = false;
} else {
    window.Strgen = Strgen;
    is_browser = true;
}; // source: http://www.matteoagosti.com/blog/2013/02/24/writing-javascript-modules-for-both-browser-and-node/
