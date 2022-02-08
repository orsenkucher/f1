rem ------------------------------------------------OUTPUT+COLLECT ABUNDANCES
set directory=.
set dirOUT=Output
set PathROOT=C:\\Users\\Orsen\\Desktop\\dev\\f1\\ma\\FitRSF.wl
@REM set PathMathematica=C:\\Program Files\\Wolfram Research\\Mathematica\\11.1\\math.exe
set PathMathematica=C:\\Program Files\\Wolfram Research\\Mathematica\\11.1\\wolframscript.exe
@REM "%PathMathematica%" -run Get[\"%PathROOT%"]
"%PathMathematica%" -script "%PathROOT%"
