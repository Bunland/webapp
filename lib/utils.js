'use strict';
// Enctoder uft8
const uft8 = new TextEncoder();
const encoder = (data) => uft8.encode(data);
export {
    encoder
}