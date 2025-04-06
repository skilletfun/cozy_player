from rest_framework import serializers

from artists.models import Artist


class ArtistSerializer(serializers.ModelSerializer):
    class Meta:
        model = Artist
        fields = "__all__"


class ArtistInfoSerializer(serializers.ModelSerializer):
    tracks_count = serializers.IntegerField()
    duration = serializers.IntegerField()

    class Meta:
        model = Artist
        fields = "__all__"
