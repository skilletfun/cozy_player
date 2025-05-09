from common.filters import BaseFilterSet
from tracks.models import Track


class TrackFilter(BaseFilterSet):
    class Meta(BaseFilterSet.Meta):
        model = Track
        fields = ["title", "artist"]
