import { ENV } from "$lib/shared.svelte";

export async function get(url, params) {
  const urlParams = new URLSearchParams(params);
  const finalUrl = `${ENV.API_URL}${url}${params ? "?" + urlParams.toString() : ""}`;
  return await fetch(finalUrl);
}

export async function post(url, data) {
  return await fetch(
    `${ENV.API_URL}${url}`, 
    {
      headers: {
        "Content-Type": "application/json",
      },
      method: "post",
      body: JSON.stringify(data),
    }
  );
}
