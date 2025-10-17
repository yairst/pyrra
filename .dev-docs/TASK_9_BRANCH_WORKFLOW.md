# Task 9: Branch Workflow Guide

## Branch Structure

**Two branches for Task 9 work:**

1. **`dev-tools-and-docs`** (current branch)
   - Has ALL files: .dev-docs, .kiro, cmd/, scripts/, prompts/, testing/
   - Use for: validation, testing, documentation reference
   - All validation tools available

2. **`add-dynamic-burn-rate`** (PR branch)
   - Clean branch: only core implementation files
   - No dev files: no .dev-docs, .kiro, cmd/, scripts/
   - Ready for upstream PR

## Workflow for Task 9

### Step 1: Work in dev-tools-and-docs (Current)

```bash
# You're already here
git branch  # Should show: * dev-tools-and-docs

# Do validation work, run tests, use validation tools
# Reference documentation in .dev-docs/
```

### Step 2: If Core Files Need Changes

If you find bugs or need to modify core implementation files (slo/, ui/, proto/, kubernetes/, etc.):

```bash
# 1. Make changes in dev-tools-and-docs branch
# Edit files: slo/rules.go, ui/src/components/*, etc.

# 2. Commit to dev-tools-and-docs
git add <modified-files>
git commit -m "fix: description of fix"

# 3. Note the commit hash
git log -1  # Copy the commit hash (e.g., abc1234)

# 4. Switch to PR branch
git checkout add-dynamic-burn-rate

# 5. Cherry-pick the fix
git cherry-pick abc1234

# 6. Verify it worked
git log -1  # Should show your fix
go test ./...  # Run tests

# 7. Push both branches
git push origin add-dynamic-burn-rate
git checkout dev-tools-and-docs
git push origin dev-tools-and-docs
```

### Step 3: Final PR Preparation (Task 9.4)

```bash
# Switch to PR branch
git checkout add-dynamic-burn-rate

# Verify it's clean (no dev files)
ls -la .dev-docs  # Should not exist
ls -la .kiro      # Should not exist
ls -la cmd/       # Should not exist

# Final verification
go build -o pyrra .
cd ui && npm run build && cd ..
go test ./...

# All good? Ready for PR!
```

## Quick Reference Commands

### Check which branch you're on
```bash
git branch
```

### Switch branches
```bash
git checkout dev-tools-and-docs      # For validation work
git checkout add-dynamic-burn-rate   # For PR preparation
```

### Cherry-pick a commit
```bash
# From dev-tools-and-docs to add-dynamic-burn-rate
git log -1                           # Get commit hash
git checkout add-dynamic-burn-rate
git cherry-pick <commit-hash>
```

### View file from other branch (without switching)
```bash
# View a file from add-dynamic-burn-rate while in dev-tools-and-docs
git show add-dynamic-burn-rate:slo/rules.go
```

## What Goes Where?

### Core Implementation Files (Both Branches)
- `slo/*.go` - Backend logic
- `ui/src/**/*` - Frontend code
- `proto/**/*` - Protobuf definitions
- `kubernetes/**/*` - Kubernetes integration
- `*.go` (root level) - Main application files
- `examples/*.yaml` - Example configurations
- Test files: `*_test.go`, `*.spec.tsx`

### Dev-Only Files (Only dev-tools-and-docs)
- `.dev-docs/**/*` - All documentation
- `.kiro/**/*` - Specs, steering, hooks
- `cmd/**/*` - Validation tools
- `scripts/**/*` - Validation scripts
- `prompts/**/*` - AI session prompts
- `testing/**/*` - Test utilities
- `Dockerfile.custom`, `Dockerfile.dev`

## Troubleshooting

### "I made changes in the wrong branch"
```bash
# If you made changes in add-dynamic-burn-rate but they should be in dev-tools-and-docs:
git stash                            # Save changes
git checkout dev-tools-and-docs
git stash pop                        # Apply changes
# Now commit and cherry-pick as normal
```

### "Cherry-pick has conflicts"
```bash
# Resolve conflicts manually
git status                           # See conflicted files
# Edit files to resolve conflicts
git add <resolved-files>
git cherry-pick --continue
```

### "I need to see what's in the PR branch"
```bash
# While in dev-tools-and-docs, view PR branch files
git diff add-dynamic-burn-rate       # See all differences
git show add-dynamic-burn-rate:slo/rules.go  # View specific file
```

## Summary

- **Work**: `dev-tools-and-docs` (has all tools and docs)
- **Fix core files**: Commit to `dev-tools-and-docs`, then cherry-pick to `add-dynamic-burn-rate`
- **PR**: From `add-dynamic-burn-rate` (clean, no dev files)
- **Preserve**: `dev-tools-and-docs` keeps all development artifacts forever
