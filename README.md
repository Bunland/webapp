#webapp

Build webapps with bun.js and chromium-based browsers.
This is a project still under construction, please be patient.

#Get Started

To get started with the project, clone the same as follows. open a console and paste this

``` 
git clone https://github.com/Bunland/webapp.git
```
Now go to the project with:
```
cd webapp
```

install dependencies (The necessary dependencies for our Svelte Ui will be installed).

```
cd myapp 
```

```
bun install
```

Run the application in development mode. 

Go back to the main "webapp" folder and type in the console:
```
bun run dev
```

Build the webapp.

Unfortunately "for now", the only way to build our webapp is using the Golang language, so make sure you have golang installed. You can follow the following link which will help you to install Golang easily:

https://github.com/golang/go/wiki/Ubuntu


Once you have golang installed write in the console in the "webapp" directory:

```
bun run build
```

Next, Golang will take care of packaging your application, you will see that a folder called "webapp_1.0.0" and a file called "webapp_1.0.0.deb" have been created, with the latter you will be able to install your application on your computer.



