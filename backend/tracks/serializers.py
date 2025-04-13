from rest_framework import serializers

from .models import Track


class TrackSerializer(serializers.ModelSerializer):
    artist_name = serializers.CharField(read_only=True)

    class Meta:
        model = Track
        fields = "__all__"
