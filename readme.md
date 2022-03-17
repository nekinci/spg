# spg 

![ci](https://github.com/nekinci/spg/actions/workflows/ci.yaml/badge.svg)
![ci](https://github.com/nekinci/spg/actions/workflows/dockerhub.yaml/badge.svg)

`spg` is a tool that generates Spring profiles based on the current environment. 

If you develop your micro services with real environments and your yaml files not up to date probably you spend time fixing.
The project aims to save you time and you can use it during development and testing.

# Installation
There are several ways to install the project.

If you are using MacOS you can install it with Homebrew:
```bash
brew tap nekinci/tap
brew install spg
```

If you want to use it with Docker you can install it with Docker:
```bash
docker pull nekinci/spg:latest
alias spg='docker run  -v ${PWD}:/go/src/app -v ${HOME}/.spg:/home/spg/.spg nekinci/spg:latest'
```

If you want to use npm you can install it with npm:
```bash
npm install -g spgjs
```

Or you can install it with manual installation of the project:

You can visit [release page](https://github.com/nekinci/spg/releases) to see the latest version. After that you can download the project and add it to your PATH.

# Usage

Firstly you need to create a configuration file. As I explained in the [Configuration](#Configuration) section you can use the configuration file with the your environment.

And then you must set the configuration file with the `spg` command:
```bash
spg config set <path to configuration file>
```

If you want to see the current configuration file you can use the `spg config print` command:

```bash
spg config print
```

Or you can unset the configuration file with the `spg config unset` command:
````bash
spg config unset
````

To generate a profile you can use the `spg generate` command. The following example shows how to generate a profile for the `my-profile` environment.
`my-profile` environment must be defined in the configuration file. `-o` or `--output` option is optional. The default value is `application-result.yaml`. If you want to change the output file you can use the `-o` or `--output` option.
And lastly you must be give the already exist profile file as an parameter.
The following command is a working example:

```bash
spg generate -p my-profile -o my-new-profile.yml application.yml application-dev.yml
```

    
#### Configuration
Example:
You have a `application.yaml` file with the following content:
    
```yaml
server:
  port: 8080
client:
  urls:
    serviceA: servicea.test.com
    serviceB: serviceb.test.com
```

Let's assume you want debug the application in prp environment. But you don't have a `application-prp.yaml` file. And you don't know prp urls. 
Probably you don't want to create a new file. Because your job is to debug or develop your application. With `spg` you can generate a `application-prp.yaml` file with the following content:

````yaml
server:
  port: 8080
client:
  urls:
    serviceA: servicea.prp.com
    serviceB: serviceb.prp.com
````


But first you need to define trainer config file that contains url information of your services. Because spg don't know the url of your services. The following config file is an example:

````yaml
version: v1
information:
  absolute-configs:
    - config-key: spring.http.encoding.force # means change the value of spring.http.encoding.force to true
      environment:
        oc: true
        test: false
        prp: true
        prod: false
    - config-key: keycloak.securityConstraints[0].authRoles[0] # means change the first role of the first security constraint
      environment:
        oc: user
        test: user
        prp: admin
        prod: admin
    - config-key: keycloak.roles[].value # means change all role values under keycloak.roles
      environment:
        oc: user
        test: user
        prp: admin
        prod: admin
  fields:
    - keys: # metadata keys for decide which fields to change
        - a-service
        - aService
      type: url
      environment:
        oc: 
          scheme: http
          value: a-service:8080
        test:
          scheme: http
          value: a-service-test.cloud.com
        prp:
          scheme: http
          value: a-service-prp.cloud.com
        prod:
          scheme: https
          value: a-service.com
````


### Known Issues

- **Don't handled yet metadata in trainer yaml file** (e.g. `information.fields.keys`)
- **Don't handled yet wildcard feature** (e.g. `absolute-configs.config-key`)
- **Config array selectors only support for primitive types. Don't supported for complex types.** (e.g. `maps or arrays`) 