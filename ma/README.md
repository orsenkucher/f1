```bash
wolframscript -script FitRSF.wl
```

```batch
rem ------------------------------------------------OUTPUT+COLLECT ABUNDANCES
rem ------------------------------------------------
set directory=.
set dirOUT=Output
set PathROOT=C:\\Work\\Oscillations\\FortranFitSimplyEXFORMathem_old3\\FitRSF2osc_new.ma\
set PathROOT=C:\\Work\\Oscillations\\FortranFitSimplyEXFORMathem_old3\\FitRSF2osc_newTSE.ma\
set PathROOT=H:\\Sasha\\work\\Oscillations\\FortranFitSimplyEXFORMathem_old3\\FitRSF2osc_newTSE.ma\
set PathROOT=H:\\Sasha\\work\\Oscillations\\FortranFitSimplyEXFORMathem_old3\\FitRSF2osc_new.ma\
set PathMathematica=C:\Program Files\Wolfram Research\Mathematica\9.0\math.exe
"%PathMathematica%" -run Get[\"%PathROOT%"]
```
