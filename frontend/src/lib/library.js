import { acts } from '@tadashi/svelte-notification';
import { API } from "$lib/api.js";

export async function rescan() {
    try {
        acts.add({mode: 'normal', message: 'Start rescan library...', lifetime: 2});
        await API.Library.rescan();
        acts.add({mode: 'success', message: 'Rescan library complete', lifetime: 2});
    } catch (e) {
        acts.add({mode: 'danger', message: `Error while rescan: ${e.message}`, lifetime: 5});
    }
}
