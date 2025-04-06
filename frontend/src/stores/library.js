export const libraryURL = "http://127.0.0.1:8000/api/library/";

export async function rescanLibrary() {
    const response = await fetch(libraryURL + 'rescan/');
}
