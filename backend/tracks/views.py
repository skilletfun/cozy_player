from django.http import FileResponse, HttpResponse
from django.db.models import F
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.generics import ListAPIView, RetrieveAPIView, RetrieveUpdateAPIView

from .models import Track
from .serializers import TrackSerializer
from .filters import TrackFilter


class TrackListAPIView(ListAPIView):
    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TrackFilter


class TrackQueueAPIView(ListAPIView):
    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer


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
                "Cache-Control": "public, max-age=604800, immutable",
                "Content-Type": "image",
            },
        )
