COVER_HEADERS: dict[str, str] = {
    "Cache-Control": "public, max-age=43200, immutable",
    "Content-Type": "image",
}

TRACK_HEADERS: dict[str, str] = {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Expose-Headers": "Content-Length, Content-Range",
    "Accept-Ranges": "bytes",
}
