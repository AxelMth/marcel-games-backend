# GitHub Actions Workflows

## Level of the Day Population

This repository contains two workflows for populating the "Level of the Day" feature:

### 1. Standard Workflow (`populate-level-of-the-day.yml`)
- Runs daily at midnight UTC
- Uses Go directly on the GitHub Actions runner
- Can be triggered manually with an optional target date

### 2. Docker Workflow (`populate-level-docker.yml`)
- Runs daily at midnight UTC
- Uses the application's Docker container for consistency with production
- Can be triggered manually with an optional target date

## Setup Requirements

Before these workflows can run successfully, you need to:

1. **Add the DATABASE_URL secret to your repository:**
   - Go to Settings > Secrets and variables > Actions
   - Click "New repository secret"
   - Name: `DATABASE_URL`
   - Value: Your PostgreSQL connection string

2. **Ensure your database has the required schema:**
   - The `LevelOfTheDay` table must exist
   - Run migrations if needed

## Manual Triggering

Both workflows can be triggered manually:

1. Go to the Actions tab in your repository
2. Select the workflow you want to run
3. Click "Run workflow"
4. Optionally provide a target date in YYYY-MM-DD format
5. Click "Run workflow" button

If no date is provided, the workflow will generate a level for tomorrow.

## How It Works

The level generation process:
1. Connects to the database using the DATABASE_URL secret
2. Checks if a level already exists for the target date
3. If not, generates random country codes using a deterministic seed based on the date
4. Creates a new LevelOfTheDay entry in the database

## Monitoring

- Both workflows will create a GitHub issue if they fail
- Check the Actions tab for workflow run history and logs
- Successful runs will log the created level details

## Choosing Between Workflows

- Use the **standard workflow** for faster execution and lower resource usage
- Use the **Docker workflow** if you want to ensure consistency with your production environment

You can disable one of the workflows if you only need one approach.
