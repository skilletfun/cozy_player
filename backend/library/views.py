import logging
import os

import music_tag
from django.conf import settings
from django.db.models import Count
from rest_framework.views import APIView
from rest_framework.response import Response

from artists.models import Artist
from common.utils import is_track_file
from tracks.models import Track


logger = logging.getLogger(__name__)


class RescanLibraryAPIView(APIView):
    def get(self, request):
        artists_qs = Artist.objects.all().prefetch_related("tracks")
        artists: dict[str, Artist] = {a.name: a for a in artists_qs}
        tracks: dict[Artist, dict[str, Track]] = {
            a: {t.path: t for t in a.tracks.all()} for a in artists.values()
        }

        logger.info(f"Rescaning music folder: {settings.MUSIC_FOLDER}")
        for artist_folder in os.listdir(settings.MUSIC_FOLDER):
            logger.info(f"Rescaning artist: {artist_folder}")
            artist_path = os.path.join(settings.MUSIC_FOLDER, artist_folder)
            if not os.path.isdir(artist_path):
                logger.info(f"Skip because is not folder: {artist_folder}")
                continue

            if not (artist := artists.get(artist_folder)):
                artist = Artist.objects.create(name=artist_folder)
                artists[artist_folder] = artist
                tracks[artist] = {}

            new_artist_tracks: list[str] = []

            for root, dirs, files in os.walk(artist_path):
                for filename in files:
                    if not is_track_file(track_name := os.path.join(root, filename)):
                        logger.info(f"Skip because is not track: {track_name}")
                        continue

                    # Если трека в списке нету, то добавим в список новых
                    # Если есть, то удалим из списка треков исполнителя (оставшиеся сотрутся из БД)
                    if not tracks[artist].pop(track_name, None):
                        new_artist_tracks.append(track_name)

            for track in tracks[artist].values():
                logger.info(f"Delete track: {track.path}")
                track.delete()

            new_tracks: list[Track] = []
            for track_path in new_artist_tracks:
                tags = music_tag.load_file(track_path)
                new_tracks.append(
                    Track(
                        title=tags["title"].first or "No title",
                        artist=artist,
                        path=track_path,
                        duration=int(tags["#length"].first or 0),
                    )
                )

            logger.info(f"Total new tracks for {artist.name}: {len(new_tracks)}")
            Track.objects.bulk_create(new_tracks)

        Artist.objects.annotate(total_tracks=Count("tracks")).filter(
            total_tracks=0
        ).delete()

        return Response()
