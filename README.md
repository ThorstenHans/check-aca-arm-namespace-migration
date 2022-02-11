# Verify Azure Container Apps ARM namespace migration

This is a simple action that verifies if any of your Azure Container Apps instances have been migrated from the ARM namespace `Microsoft.Web` to `Microsoft.App`.

If at least one Azure Container App is migrated, the action will fail with a non-zero exit code.

## Usage

```yaml
name: 'Check'
on:
  workflow_dispatch:
  schedule:
    - cron: '0 */12 * * *'
jobs:
  check:
    name: Verify Azure Container Apps namespace migration
    runs-on: ubuntu-latest
    steps:
      - name: Verify
        id: verify
        uses: ThorstenHans/check-aca-arm-namespace-migration@v1
        with:
          azure-client-id: ${{ secrets.AZURE_CLIENT_ID}}
          azure-client-secret: ${{ secrets.AZURE_CLIENT_SECRET }}
          azure-subscription-id: ${{ secrets.AZURE_SUB_ID }}
          azure-tenant-id: ${{ secrets.AZURE_TENANT_ID }}
```
