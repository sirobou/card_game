import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
// import { VitePluginFonts } from "vite-plugin-fonts"
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // VitePluginFonts({
    //   google: {
    //     families: [],
    //   },
    // }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
})
