import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    lib: {
      // entry: 'src/main.js',
      entry: 'index.html',
      formats: ['cjs'],
      name: 'app',
      fileName: () => 'bundle.js'
    },
    outDir: '../www'
  }
})
