import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: vitePreprocess(),
	kit: {
		adapter: adapter({
      pages: 'build',  // Куда складывать файлы (по умолчанию)
      assets: 'build',
      fallback: 'index.html', // Для SPA-режима
    }),
		prerender: {
		  entries: ["*"],
		}
	}
};

export default config;
