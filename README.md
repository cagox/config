# github.com/cagox/config

This simple package is the result of my search to find a simple standardize way to handle configuration files for my Go project.

It operations by loading the contents of a JSON file into a struct that you define. See github.com/cagox/config/main for an example of how this works.

The steps to use this package amount to the following:

1. Create your own local configuration struct (or use ours if you think it works).
2. Embed our ConfigurationStruct in your struct.
3. Create a JSON configuration file for your project, and set an environmental variable to point to it's location.
4. Call LoadConfigs()
5. Reference the configuration struct for configuration variables.  
