// Dependencies
import { dlopen, suffix, FFIType } from "bun:ffi";
// Get SO Path
const path = `./webapp.${suffix}`;
const location = new URL(path, import.meta.url).pathname
// Get Symbols/Functions from So
const { symbols } = dlopen(location, {
    CreateWindow: {
        args: [FFIType.ptr]
    }
})

export {
    symbols
}