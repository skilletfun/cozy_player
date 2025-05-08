from collections import defaultdict


def shuffle_tracks(tracks: dict[int, list[dict]]) -> list[dict]:
    result: list[dict] = []
    order: list[int] = sorted(tracks.keys(), key=lambda x: len(tracks[x]), reverse=True)

    while tracks:
        if len(tracks) == 1:
            result.extend(list(tracks.values())[0])
            break

        result.append(tracks[order[0]].pop())
        if not tracks[order[0]]:
            tracks.pop(order.pop(0))

        order = sorted(order[1:], key=lambda x: len(tracks[x]), reverse=True) + order[:1]

    return result


def shuffle_queue(tracks: list[dict]) -> list[dict]:
    parts: dict[int, dict[int, list[dict]]] = defaultdict(lambda: defaultdict(list))

    for track in tracks:
        parts[track["play_count"]][track["artist"]].append(track)

    result: list[dict] = []
    for c in sorted(parts.keys()):
        result.extend(shuffle_tracks(parts[c]))

    return result
