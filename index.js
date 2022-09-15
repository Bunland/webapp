// Dependencies
import { webApp } from "./lib/webapp";
// Instantiate a new webapp
const webapp = new webApp();
// Options with which the webapp will start
const options = {
  /*
  The Vite Server starts on the port 5173
  so let's tell the webapp to start at the url "http://localhost:5173"
  */
  Url: `http://localhost:5173/`, 
  Size: "800,600", // Width, Height
  FullScreen: false,
  /*
  Add an array of flags, 
  full list of chromium flags here: https://peter.sh/experiments/chromium-command-line-switches/
  Please dont remove --class, 
  You can change the class name as you like. example "--class=MyAppName"
  */
  Args: ["--class=WebApp"]
};
// Save the configuration for when the application is built.
await Bun.write("build.json", JSON.stringify(options));
// Start webapp
await webapp.createWindow(options);
