from django.db.models import Count, Sum
from django.http import FileResponse
from rest_framework.generics import ListAPIView, RetrieveAPIView
from django_filters.rest_framework import DjangoFilterBackend

from common.headers import COVER_HEADERS
from .models import Artist
from .serializers import ArtistSerializer, ArtistInfoSerializer
from .filters import ArtistFilter


class ArtistListAPIView(ListAPIView):
    """Список исполнителей"""

    queryset = Artist.objects.all()
    serializer_class = ArtistSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = ArtistFilter
    pagination_class = None


class ArtistGetAPIView(RetrieveAPIView):
    """Получить исполнителя"""

    queryset = Artist.objects.annotate(
        tracks_count=Count("tracks"),
        duration=Sum("tracks__duration"),
    )
    serializer_class = ArtistInfoSerializer


class ArtistCoverGetAPIView(RetrieveAPIView):
    """Получить обложку исполнителя"""

    queryset = Artist.objects.all()

    def retrieve(self, request, *args, **kwargs):
        artist: Artist = self.get_object()
        return FileResponse(open(artist.cover, "rb"), headers=COVER_HEADERS)
