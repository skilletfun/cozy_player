import os
from django.http import FileResponse, HttpResponse
from django.db.models import F
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.generics import ListAPIView, RetrieveAPIView, RetrieveUpdateAPIView
from rest_framework.response import Response
from rest_framework.exceptions import NotFound

from common.headers import COVER_HEADERS, TRACK_HEADERS
from .models import Track
from .serializers import TrackSerializer
from .filters import TrackFilter
from .utils import shuffle_queue


class TrackListAPIView(ListAPIView):
    """Отдает список треков"""

    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TrackFilter
    pagination_class = None


class TrackQueueAPIView(ListAPIView):
    """Генерирует очередь воспроизведения"""

    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer
    pagination_class = None

    def list(self, request, *args, **kwargs):
        tracks = self.get_queryset()
        serializer = self.get_serializer(tracks, many=True)
        return Response(shuffle_queue(serializer.data))


class TrackGetUpdateAPIView(RetrieveUpdateAPIView):
    """PATCH / PUT для обновления данных трека, GET для получения файла на воспроизведение"""

    queryset = Track.objects.all()
    serializer_class = TrackSerializer

    def retrieve(self, request, *args, **kwargs):
        if not os.path.exists(track_path := self.get_object().path):
            raise NotFound(f"Track not found: {track_path}")
        return FileResponse(open(track_path, "rb"), headers=TRACK_HEADERS)


class TrackCoverGetAPIView(RetrieveAPIView):
    """Отдает обложку для трека"""

    queryset = Track.objects.all()

    def retrieve(self, request, *args, **kwargs):
        return HttpResponse(self.get_object().cover, headers=COVER_HEADERS)
