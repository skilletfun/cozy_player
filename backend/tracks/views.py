from django.http import FileResponse, HttpResponse
from django.db.models import F
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.generics import ListAPIView, RetrieveAPIView, RetrieveUpdateAPIView
from rest_framework.response import Response

from .models import Track
from .serializers import TrackSerializer
from .filters import TrackFilter
from .utils import shuffle_queue


class TrackListAPIView(ListAPIView):
    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TrackFilter


class TrackQueueAPIView(ListAPIView):
    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer

    def list(self, request, *args, **kwargs):
        tracks = self.get_queryset()
        serializer = self.get_serializer(tracks, many=True)
        return Response(shuffle_queue(serializer.data))


class TrackGetUpdateAPIView(RetrieveUpdateAPIView):
    queryset = Track.objects.all()
    serializer_class = TrackSerializer

    def retrieve(self, request, *args, **kwargs):
        return FileResponse(
            open(self.get_object().path, "rb"),
            headers={
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Expose-Headers": "Content-Length, Content-Range",
                "Accept-Ranges": "bytes",
            },
        )


class TrackCoverGetAPIView(RetrieveAPIView):
    queryset = Track.objects.all()

    def retrieve(self, request, *args, **kwargs):
        return HttpResponse(
            self.get_object().cover,
            headers={
                "Cache-Control": "public, max-age=43200, immutable",
                "Content-Type": "image",
            },
        )
