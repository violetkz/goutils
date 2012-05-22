goutils
=======

the utilities with Go language

config.go:
    the simple code for parsing config file.
    
    Example:
        cf, err = config("xxx.cfg")
        if err !=nil {
            // Open config file failed.
        }
        
        cf.parse()
        
        value := cf.get("key")
        fmt.Printf("%s\n", value)
                
