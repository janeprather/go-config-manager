# go-config-manager

After making over a half dozen apps that all used the same bits of code to
read configuration data out of a json file, I decided to make this package
of it and stop rewriting the entire wheel.

# Usage

Import it.

```
import (
  ...

  cfgMgr "github.com/janeprather/go-config-manager"
)
```

Design your own configuration struct.  This can use json tags to
customize the json fields if desired.

```
type MyConfig struct {
  URL       string
  AuthToken string
  IsActive  bool `json:"is_active"`
}
```

Then instantiate one of your config items.

```
config := &MyConfig{}
```

Determine a config filename somehow.

```
fileName := "/path/to/config.file"
```

Finally, load the configuration.

```
err := cfgMgr.LoadConfig(config, fileName)
if err != nil {
  fmt.Printf("error: %s", err.Error())
}
```

For the above to work, your config file should something like:

```
{
  "URL": "http://some.url/path",
  "AuthToken": "12345",
  "is_active": true
}
```

You can generate/edit a config struct and save it as well.

```
config := &MyConfig{}

// Only save if the config file does not already exist  
err := cfgMgr.SaveNewConfig(config, fileName)

// Save, overwriting any existing config
err := cfgMgr.SaveConfig(config, fileName)
```

Note that you can nest structs for more complex configuration designs:

```
type MyConfig struct {
  AppName    string
  AppVersion string
  ServiceOne struct {
    User     string
    Password string
  }
  ServiceTwo struct {
    Site      string
    AuthToken string
  }
}
```

And you can break it down into separate smaller structs.

```
type ServiceOneConfig struct {
  User     string
  Password string
}

type ServiceTwoConfig struct {
  Site      string
  AuthToken string
}

type MyConfig {
  AppName string
  AppVersion string
  ServiceOne *ServiceOneConfig
  ServiceTwo *ServiceTwoConfig
}
```

Note that if you wish to use SaveConfig(...), and are using pointers as above,
then you will need to instantiate each struct for which you wish to have fields
saved.

```
config := &MyConfig{}
config.ServiceOne = &ServiceOneConfig{}
config.ServiceTwo = &ServiceTwoConfig{}

// now we should be able to save an empty
// config with all the nested fields

cfgMgr.SaveConfig(config, fileName)
```
