#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define GCP Region
REGION="asia-southeast1"
SERVICE_NAME="demo-bion"

echo "🧹 Starting Google Cloud infrastructure cleanup for '$SERVICE_NAME' in '$REGION'..."

# 1. Delete the Cloud Run service
echo "========================================="
echo "🗑️ Deleting Cloud Run service: $SERVICE_NAME"
echo "========================================="
gcloud run services delete "$SERVICE_NAME" --region="$REGION" --quiet || echo "⚠️ Service already deleted or not found."

# 2. Delete the Artifact Registry repository
echo ""
echo "========================================="
echo "📦 Deleting Artifact Registry repository: $SERVICE_NAME"
echo "========================================="
gcloud artifacts repositories delete "$SERVICE_NAME" --location="$REGION" --quiet || echo "⚠️ Repository already deleted or not found."

echo ""
echo "✅ Cleanup complete!"
