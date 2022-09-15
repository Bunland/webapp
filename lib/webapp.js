'use strict';
// Dependencies
import { ptr } from "bun:ffi";
import { symbols } from "./ffi";
import { encoder } from "./utils";
// Class webAPp
class webApp {
    createWindow (data) {
        return symbols.CreateWindow(ptr(encoder(JSON.stringify(data))))
    }
}

export {
    webApp
}