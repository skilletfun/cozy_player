from django.http import FileResponse
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.generics import ListAPIView, RetrieveAPIView

from .models import Track
from .serializers import TrackSerializer
from .filters import TrackFilter


class TrackListAPIView(ListAPIView):
    queryset = Track.objects.all()
    serializer_class = TrackSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TrackFilter


# class TrackCoverGetAPIView(RetrieveAPIView):
#     queryset = Track.objects.all()
#
#     def retrieve(self, request, *args, **kwargs):
#         track: Track = self.get_object()
#         return FileResponse(open(artist.cover, "rb"))
