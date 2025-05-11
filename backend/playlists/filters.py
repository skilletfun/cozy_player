from common.filters import BaseFilterSet
from .models import Playlist


class PlaylistFilter(BaseFilterSet):
    class Meta(BaseFilterSet.Meta):
        model = Playlist
        fields = ["name"]
