

Simple go application to generate config files from environment varialbles (key value pairs) based on specific naming convention. Used in docker container to generate the configuration files based on the environment variables.

The environment should be in form:

```
FILENAME.EXT_any.key=value
```

The extracted key value pairs will be transformed to other formats based on the extension.

Supported formats/extensions:

 * xml: hadoop xml format
 * env, cfg: `key=value` format
 * conf, sh: `export key=value` format
 * properties: `key: properties`


