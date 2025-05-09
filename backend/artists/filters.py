from common.filters import BaseFilterSet
from artists.models import Artist


class ArtistFilter(BaseFilterSet):
    class Meta(BaseFilterSet.Meta):
        model = Artist
        fields = ["name"]
