# Task 9.3: Interactive UI Testing Guide

## Quick Start

**Automated tests are complete and passed!** ✅

Now we need you to perform quick interactive UI testing to verify the user interface works correctly.

## Prerequisites ✅

- ✅ Prometheus running on port 9090
- ✅ Pyrra API running on port 9099
- ✅ Pyrra Backend running on port 9444
- ✅ 26 SLOs deployed in cluster
- ✅ Recording rules generating data

## Testing Steps

### Step 1: Open Pyrra UI

**Option A - Development UI** (Recommended for testing):
```bash
cd ui
npm start
# Opens http://localhost:3000
```

**Option B - Production UI** (If already built):
```bash
# Just open http://localhost:9099 in browser
```

### Step 2: List Page Validation (2 minutes)

**Open**: http://localhost:3000 (or http://localhost:9099)

**Check**:
1. [ ] Page loads successfully (no errors)
2. [ ] "Burn Rate" column is visible
3. [ ] See gray "Static" badges with lock icon
4. [ ] See green "Dynamic" badges with eye icon
5. [ ] Count SLOs: Should see ~26 SLOs total
6. [ ] Hover over badges - tooltips appear
7. [ ] No console errors (press F12 to check)

**Expected**:
- List page displays all SLOs
- Burn rate badges show correctly
- Tooltips provide context
- No JavaScript errors in console

### Step 3: Detail Page - Ratio Indicator (2 minutes)

**Click on**: "test-dynamic-apiserver" (dynamic ratio SLO)

**Check**:
1. [ ] Detail page loads successfully
2. [ ] Burn rate type badge in header (green "Dynamic")
3. [ ] Alerts table shows threshold values (not "Traffic-Aware")
4. [ ] Threshold values are numbers (e.g., 0.140, 0.070)
5. [ ] "Error Budget Consumption" column shows percentages
6. [ ] Hover over threshold - tooltip shows traffic context
7. [ ] Graphs load correctly
8. [ ] No console errors

**Expected**:
- Threshold calculations display correctly
- Enhanced tooltips work
- All graphs render
- No crashes or errors

### Step 4: Detail Page - Latency Indicator (2 minutes)

**Click on**: "test-latency-dynamic" (dynamic latency SLO)

**Check**:
1. [ ] Detail page loads successfully
2. [ ] Threshold values display correctly
3. [ ] Latency-specific graphs render
4. [ ] Tooltips work correctly
5. [ ] No console errors

**Expected**:
- Latency indicator works like ratio
- Threshold calculations correct
- No errors

### Step 5: Missing Metrics Error Handling (2 minutes)

**Click on**: "test-missing-metrics-dynamic" (fictional metrics)

**Check**:
1. [ ] Detail page loads (NO WHITE PAGE CRASH!)
2. [ ] Tiles show "No data" or appropriate message
3. [ ] Threshold column shows fallback (not crash)
4. [ ] Click burn rate graph button - NO CRASH
5. [ ] Console may show warnings (expected)
6. [ ] Page remains functional

**Expected**:
- Graceful error handling
- No crashes
- User-friendly error messages
- Page remains usable

### Step 6: BoolGauge Indicator (1 minute)

**Click on**: "test-bool-gauge-dynamic" (dynamic bool gauge)

**Check**:
1. [ ] Detail page loads successfully
2. [ ] Threshold values display correctly
3. [ ] No console errors

**Expected**:
- BoolGauge works correctly
- Threshold calculations display
- No errors

### Step 7: Static SLO Comparison (1 minute)

**Click on**: "test-static-apiserver" (static ratio SLO)

**Check**:
1. [ ] Detail page loads successfully
2. [ ] Gray "Static" badge in header
3. [ ] Alerts table shows "Factor" column (not "Error Budget Consumption")
4. [ ] Threshold values are static (e.g., 0.700, 0.350)
5. [ ] No dynamic threshold calculations
6. [ ] No console errors

**Expected**:
- Static SLOs work unchanged
- No regressions
- Different UI from dynamic SLOs

## Quick Validation Checklist

### Must Pass ✅
- [ ] List page loads with all SLOs
- [ ] Burn rate badges display correctly
- [ ] Detail pages load for all indicator types
- [ ] Threshold calculations work
- [ ] Missing metrics don't crash page
- [ ] No critical console errors

### Nice to Have ✅
- [ ] Enhanced tooltips work
- [ ] Graphs render correctly
- [ ] Performance feels responsive
- [ ] No visual glitches

## What to Report

### If Everything Works ✅
Just say: **"All UI tests passed"** or **"UI validation complete"**

### If Issues Found ❌
Report:
1. Which SLO/page had the issue
2. What you expected to see
3. What actually happened
4. Any console errors (copy/paste)
5. Screenshot if helpful

## Expected Time

**Total Testing Time**: ~10 minutes

- Step 1: 1 minute (open UI)
- Step 2: 2 minutes (list page)
- Step 3: 2 minutes (ratio detail)
- Step 4: 2 minutes (latency detail)
- Step 5: 2 minutes (missing metrics)
- Step 6: 1 minute (bool gauge)
- Step 7: 1 minute (static SLO)

## Notes

- **Console warnings are okay** - We expect some warnings for missing metrics
- **Focus on critical functionality** - Don't worry about minor cosmetic issues
- **No need to test every SLO** - Sample testing is sufficient
- **Previous testing was comprehensive** - This is just final confirmation

## After Testing

Once you confirm UI testing passes, we'll:
1. Mark Task 9.3 as complete
2. Update documentation
3. Proceed to Task 9.4 (Prepare for upstream submission)

## Questions?

If anything is unclear or you encounter issues, just let me know!
