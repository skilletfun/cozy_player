import django_filters

from tracks.models import Track


class TrackFilter(django_filters.FilterSet):
    class Meta:
        model = Track
        fields = ["title", "artist"]
