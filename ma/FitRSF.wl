(* ::Package:: *)

Print["Hello"];
MainDirecroty = Directory[];
Filename = "group.dat";
DataEXP = ReadList[FileNameJoin[{MainDirecroty, Filename}], Number];

Ncolumbs = 3;
NmaxpointsEXP = Length[DataEXP]/Ncolumbs - 1;
Xlabel1EXP = Table[DataEXP[[3*i + 1]], {i, 0, NmaxpointsEXP}];
Ylabel1EXP = Table[DataEXP[[3*i + 2]], {i, 0, NmaxpointsEXP}];
DYlabel1EXP = Table[DataEXP[[3*i + 3]], {i, 0, NmaxpointsEXP}];
Clear[DataExp];

(* ========== *)
KEYGAM = 2; 

(* 
KEYGAM= 0 G-Const, 1 G=SMLO, 2 G=EGLO
*)

KEYGAMf = 2; 

(*
KEYGAMf= 1 G-Const, 2 G=G(E)-GDR, 3 G=G(E)-GDR,PDR, 4 G=G(E)-PDR 
*)

Ncut = 1;

Nmaxpointstemp = Length[Ylabel1EXP] - Ncut;

FordY1aver = Table[1, {i, 1, Nmaxpointstemp + 1}];

FordY1aver = Table[DYlabel1EXP[[i]], {i, 1, Nmaxpointstemp + 1}];

Data1 = Table[{Xlabel1EXP[[i]], Ylabel1EXP[[i]]}, {i, 1, 

    Nmaxpointstemp + 1}];

Filename = "KEYMODEL_ma.DAT";

KEYjoin = 

  ReadList[FileNameJoin[{MainDirecroty, Filename}], Number][[1]];

Filename = "KEYDATA_ma.DAT";

KEYdata = 

  ReadList[FileNameJoin[{MainDirecroty, Filename}], Number][[1]];

Filename = "Inputpar_ma.dat";

AlphaparIN = ReadList[FileNameJoin[{MainDirecroty, Filename}], Number];

FConst = 8.674*10^(-8);

FunctE0[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, 

   gamma12_] = -Im[(z1^2 + (I gamma12 omega z1 z2)/(I (gamma12 + 

              gamma2) omega - omega^2 + omega2^2))/(I (gamma1 + 

           gamma12) omega - omega^2 + 

        omega1^2 + (gamma12^2 omega^2)/(I (gamma12 + gamma2) omega - 

           omega^2 + 

           omega2^2)) + ((I gamma12 omega z1 z2)/(I (gamma1 + 

              gamma12) omega - omega^2 + omega1^2) + 

        z2^2)/(I (gamma12 + gamma2) omega - 

        omega^2 + (gamma12^2 omega^2)/(I (gamma1 + gamma12) omega - 

           omega^2 + omega1^2) + omega2^2)];

GammaCONST[omega_, gammarx_, omegarx_] = gammarx;

GammaSMLO[omega_, gammarx_, omegarx_] = gammarx*omega/omegarx;

GammaEGLO[omega_, gammarx_, omegarx_] = gammarx*omega^2/(omegarx^2);

If[KEYGAM == 0, GammaF[omega_, gammarx_, omegarx_] = GammaCONST[omega, gammarx, omegarx]];

If[KEYGAM == 1, GammaF[omega_, gammarx_, omegarx_]=GammaSMLO[omega, gammarx, omegarx]];

If[KEYGAM == 2, GammaF[omega_, gammarx_, omegarx_]=GammaEGLO[omega, gammarx, omegarx]];

If[KEYGAMf == 1, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, gamma1, omega2, z2, gamma2, gamma12]];

If[KEYGAMf == 2, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, GammaF[omega, gamma1, omega1], omega2, z2, gamma2, gamma12]];

If[KEYGAMf == 3, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, GammaF[omega, gamma1, omega1], omega2, z2, GammaF[omega, gamma2, omega2], gamma12]];

If[KEYGAMf == 4, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, gamma1, omega2, z2, GammaF[omega, gamma2, omega2], gamma12]];

If[KEYdata == 1 || KEYdata == 5 || KEYdata == 6, 

  SigmaConst = 1.15291*10^(7), SigmaConst = 1.0];

If[KEYdata == 1 || KEYdata == 5 || KEYdata == 6, 

  FunctE[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, 

    gamma12_] = 

   omega*FunctE1[omega, omega1, z1, gamma1, omega2, z2, gamma2, 

     gamma12]*FConst*SigmaConst, 

  FunctE[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, 

    gamma12_] = 

   FunctE1[omega, omega1, z1, gamma1, omega2, z2, gamma2, gamma12]*

    FConst*SigmaConst];

Which[KEYjoin == 

   1, {model = 

    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, 

     Gamma12p], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, 

      Z20 > 0, Gamma2p > 0, 

      Gamma12p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, 

       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, 

       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, 

       AlphaparIN[[6]]}, {Gamma12p, AlphaparIN[[7]]}}, omega, 

     Weights -> 1/FordY1aver^2], 

   Par1 = omega1p /. FITtemp["BestFitParameters"], 

   Par2 = Z10 /. FITtemp["BestFitParameters"], 

   Par3 = Gamma1p /. FITtemp["BestFitParameters"], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], 

   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, 

  KEYjoin == 

   0, {model = 

    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, 0], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, 

      Z20 > 0, 

      Gamma2p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, 

       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, 

       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, 

       AlphaparIN[[6]]}}, omega, Weights -> 1/FordY1aver^2], 

   Par1 = omega1p /. FITtemp["BestFitParameters"], 

   Par2 = Z10 /. FITtemp["BestFitParameters"], 

   Par3 = Gamma1p /. FITtemp["BestFitParameters"], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, 

  KEYjoin == 2, {model = FunctE[omega, omega1p, Z10, 0, 0, 0, 0, 0], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega1p > 0, Z10 > 0, 

      Gamma1p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, 

       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}}, omega, 

     Weights -> 1/FordY1aver^2], 

   Par1 = omega1p /. FITtemp["BestFitParameters"], 

   Par2 = Z10 /. FITtemp["BestFitParameters"], 

   Par3 = Gamma1p /. FITtemp["BestFitParameters"], Par4 = 0, Par5 = 0,

    Par6 = 0, Par7 = 0}, 

  KEYjoin == 

   3, {model = 

    FunctE[omega, AlphaparIN[[1]], AlphaparIN[[2]], AlphaparIN[[3]], 

     omega2p, Z20, Gamma2p, Gamma12p], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega2p > 0, Z20 > 0, Gamma2p > 0, 

      Gamma12p > 0}, {{omega2p, AlphaparIN[[4]]}, {Z20, 

       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, 

       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], 

   Par1 = AlphaparIN[[1]], Par2 = AlphaparIN[[2]], 

   Par3 = AlphaparIN[[3]], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], 

   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, 

  KEYjoin == 

   4, {model = 

    FunctE[omega, AlphaparIN[[1]], AlphaparIN[[2]], AlphaparIN[[3]], 

     omega2p, Z20, Gamma2p, 0], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega2p > 0, Z20 > 0, 

      Gamma2p > 0}, {{omega2p, AlphaparIN[[4]]}, {Z20, 

       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}}, omega, 

     Weights -> 1/FordY1aver^2], Par1 = AlphaparIN[[1]], 

   Par2 = AlphaparIN[[2]], Par3 = AlphaparIN[[3]], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, 

  KEYjoin == 

   5, {model = 

    FunctE[omega, AlphaparIN[[1]], Z10, Gamma1p, omega2p, Z20, 

     Gamma2p, Gamma12p], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, Z10 > 0, Gamma1p > 0, omega2p > 0, Z20 > 0, 

      Gamma2p > 0, 

      Gamma12p > 0}, {{Z10, AlphaparIN[[2]]}, {Gamma1p, 

       AlphaparIN[[3]]}, {omega2p, AlphaparIN[[4]]}, {Z20, 

       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, 

       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], 

   Par1 = AlphaparIN[[1]], Par2 = Z10 /. FITtemp["BestFitParameters"],

    Par3 = Gamma1p /. FITtemp["BestFitParameters"], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], 

   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, 

  KEYjoin == 

   6, {model = 

    FunctE[omega, AlphaparIN[[1]], Z10, Gamma1p, omega2p, Z20, 

     Gamma2p, Gamma12p], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, Z10 > 0, Gamma1p > 0, omega2p > 0, Z20 > 0, 

      Gamma2p > 0, 

      Gamma12p > 0}, {{Z10, AlphaparIN[[2]]}, {Gamma1p, 

       AlphaparIN[[3]]}, {omega2p, AlphaparIN[[4]]}, {Z20, 

       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, 

       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], 

   Par1 = AlphaparIN[[1]], Par2 = Z10 /. FITtemp["BestFitParameters"],

    Par3 = Gamma1p /. FITtemp["BestFitParameters"], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, 

  KEYjoin == 

   7, {model = 

    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, 0], 

   FITtemp = 

    NonlinearModelFit[

     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, 

      Z20 > 0, 

      Gamma2p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, 

       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, 

       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, 

       AlphaparIN[[6]]}}, omega, Weights -> 1/FordY1aver^2], 

   Par1 = omega1p /. FITtemp["BestFitParameters"], 

   Par2 = Z10 /. FITtemp["BestFitParameters"], 

   Par3 = Gamma1p /. FITtemp["BestFitParameters"], 

   Par4 = omega2p /. FITtemp["BestFitParameters"], 

   Par5 = Z20 /. FITtemp["BestFitParameters"], 

   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}];

OUTData ={{Par1},{Par2},{Par3},{Par4},{Par5},{Par6},{Par7}}; 

Export["output.dat", OUTData , "Table"];

Quit[] 


(* ::Input:: *)
(*MainDirecroty = Directory[];*)
(*Print[MainDirecroty];*)
(*(* MainDirecroty = NotebookDirectory[]; *)*)
(**)
(*Filename = "exp.dat";*)
(**)
(*DataEXP = ReadList[FileNameJoin[{MainDirecroty, Filename}], Number];*)
(**)
(*Ncolumbs = 3;*)
(**)
(*NmaxpointsEXP = Length[DataEXP]/Ncolumbs - 1;*)
(**)
(*Xlabel1EXP = Table[DataEXP[[3*i + 1]], {i, 0, NmaxpointsEXP}];*)
(**)
(*Ylabel1EXP = Table[DataEXP[[3*i + 2]], {i, 0, NmaxpointsEXP}];*)
(**)
(*DYlabel1EXP = Table[DataEXP[[3*i + 3]], {i, 0, NmaxpointsEXP}];*)
(**)
(*Clear[DataExp];*)
(**)
(**)
(**)
(*KEYGAM = 2; *)
(**)
(*(* *)
(**)
(*KEYGAM= 0 G-Const, 1 G=SMLO, 2 G=EGLO*)
(**)
(**)*)
(**)
(**)
(**)
(*KEYGAMf = 2; *)
(**)
(*(**)
(**)
(*KEYGAMf= 1 G-Const, 2 G=G(E)-GDR, 3 G=G(E)-GDR,PDR, 4 G=G(E)-PDR *)
(**)
(**)*)
(**)
(*Ncut = 1;*)
(**)
(*Nmaxpointstemp = Length[Ylabel1EXP] - Ncut;*)
(**)
(*FordY1aver = Table[1, {i, 1, Nmaxpointstemp + 1}];*)
(**)
(*FordY1aver = Table[DYlabel1EXP[[i]], {i, 1, Nmaxpointstemp + 1}];*)
(**)
(*Data1 = Table[{Xlabel1EXP[[i]], Ylabel1EXP[[i]]}, {i, 1, *)
(**)
(*    Nmaxpointstemp + 1}];*)
(**)
(*Filename = "KEYMODEL_ma.DAT";*)
(**)
(*KEYjoin = *)
(**)
(*  ReadList[FileNameJoin[{MainDirecroty, Filename}], Number][[1]];*)
(**)
(*Filename = "KEYDATA_ma.DAT";*)
(**)
(*KEYdata = *)
(**)
(*  ReadList[FileNameJoin[{MainDirecroty, Filename}], Number][[1]];*)
(**)
(*Filename = "Inputpar_ma.dat";*)
(**)
(*AlphaparIN = ReadList[FileNameJoin[{MainDirecroty, Filename}], Number];*)
(**)
(*FConst = 8.674*10^(-8);*)
(**)
(*FunctE0[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, *)
(**)
(*   gamma12_] = -Im[(z1^2 + (I gamma12 omega z1 z2)/(I (gamma12 + *)
(**)
(*              gamma2) omega - omega^2 + omega2^2))/(I (gamma1 + *)
(**)
(*           gamma12) omega - omega^2 + *)
(**)
(*        omega1^2 + (gamma12^2 omega^2)/(I (gamma12 + gamma2) omega - *)
(**)
(*           omega^2 + *)
(**)
(*           omega2^2)) + ((I gamma12 omega z1 z2)/(I (gamma1 + *)
(**)
(*              gamma12) omega - omega^2 + omega1^2) + *)
(**)
(*        z2^2)/(I (gamma12 + gamma2) omega - *)
(**)
(*        omega^2 + (gamma12^2 omega^2)/(I (gamma1 + gamma12) omega - *)
(**)
(*           omega^2 + omega1^2) + omega2^2)];*)
(**)
(*GammaCONST[omega_, gammarx_, omegarx_] = gammarx;*)
(**)
(*GammaSMLO[omega_, gammarx_, omegarx_] = gammarx*omega/omegarx;*)
(**)
(*GammaEGLO[omega_, gammarx_, omegarx_] = gammarx*omega^2/(omegarx^2);*)
(**)
(*If[KEYGAM == 0, GammaF[omega_, gammarx_, omegarx_] = GammaCONST[omega, gammarx, omegarx]];*)
(**)
(*If[KEYGAM == 1, GammaF[omega_, gammarx_, omegarx_]=GammaSMLO[omega, gammarx, omegarx]];*)
(**)
(*If[KEYGAM == 2, GammaF[omega_, gammarx_, omegarx_]=GammaEGLO[omega, gammarx, omegarx]];*)
(**)
(*If[KEYGAMf == 1, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, gamma1, omega2, z2, gamma2, gamma12]];*)
(**)
(*If[KEYGAMf == 2, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, GammaF[omega, gamma1, omega1], omega2, z2, gamma2, gamma12]];*)
(**)
(*If[KEYGAMf == 3, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, GammaF[omega, gamma1, omega1], omega2, z2, GammaF[omega, gamma2, omega2], gamma12]];*)
(**)
(*If[KEYGAMf == 4, FunctE1[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, gamma12_]=FunctE0[omega, omega1, z1, gamma1, omega2, z2, GammaF[omega, gamma2, omega2], gamma12]];*)
(**)
(*If[KEYdata == 1 || KEYdata == 5 || KEYdata == 6, *)
(**)
(*  SigmaConst = 1.15291*10^(7), SigmaConst = 1.0];*)
(**)
(*If[KEYdata == 1 || KEYdata == 5 || KEYdata == 6, *)
(**)
(*  FunctE[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, *)
(**)
(*    gamma12_] = *)
(**)
(*   omega*FunctE1[omega, omega1, z1, gamma1, omega2, z2, gamma2, *)
(**)
(*     gamma12]*FConst*SigmaConst, *)
(**)
(*  FunctE[omega_, omega1_, z1_, gamma1_, omega2_, z2_, gamma2_, *)
(**)
(*    gamma12_] = *)
(**)
(*   FunctE1[omega, omega1, z1, gamma1, omega2, z2, gamma2, gamma12]**)
(**)
(*    FConst*SigmaConst];*)
(**)
(*Which[KEYjoin == *)
(**)
(*   1, {model = *)
(**)
(*    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, *)
(**)
(*     Gamma12p], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, *)
(**)
(*      Z20 > 0, Gamma2p > 0, *)
(**)
(*      Gamma12p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, *)
(**)
(*       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, *)
(**)
(*       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, *)
(**)
(*       AlphaparIN[[6]]}, {Gamma12p, AlphaparIN[[7]]}}, omega, *)
(**)
(*     Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = omega1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par2 = Z10 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par3 = Gamma1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, *)
(**)
(*  KEYjoin == *)
(**)
(*   0, {model = *)
(**)
(*    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, 0], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, *)
(**)
(*      Z20 > 0, *)
(**)
(*      Gamma2p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, *)
(**)
(*       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, *)
(**)
(*       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, *)
(**)
(*       AlphaparIN[[6]]}}, omega, Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = omega1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par2 = Z10 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par3 = Gamma1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, *)
(**)
(*  KEYjoin == 2, {model = FunctE[omega, omega1p, Z10, 0, 0, 0, 0, 0], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega1p > 0, Z10 > 0, *)
(**)
(*      Gamma1p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, *)
(**)
(*       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}}, omega, *)
(**)
(*     Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = omega1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par2 = Z10 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par3 = Gamma1p /. FITtemp["BestFitParameters"], Par4 = 0, Par5 = 0,*)
(**)
(*    Par6 = 0, Par7 = 0}, *)
(**)
(*  KEYjoin == *)
(**)
(*   3, {model = *)
(**)
(*    FunctE[omega, AlphaparIN[[1]], AlphaparIN[[2]], AlphaparIN[[3]], *)
(**)
(*     omega2p, Z20, Gamma2p, Gamma12p], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega2p > 0, Z20 > 0, Gamma2p > 0, *)
(**)
(*      Gamma12p > 0}, {{omega2p, AlphaparIN[[4]]}, {Z20, *)
(**)
(*       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, *)
(**)
(*       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = AlphaparIN[[1]], Par2 = AlphaparIN[[2]], *)
(**)
(*   Par3 = AlphaparIN[[3]], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, *)
(**)
(*  KEYjoin == *)
(**)
(*   4, {model = *)
(**)
(*    FunctE[omega, AlphaparIN[[1]], AlphaparIN[[2]], AlphaparIN[[3]], *)
(**)
(*     omega2p, Z20, Gamma2p, 0], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega2p > 0, Z20 > 0, *)
(**)
(*      Gamma2p > 0}, {{omega2p, AlphaparIN[[4]]}, {Z20, *)
(**)
(*       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}}, omega, *)
(**)
(*     Weights -> 1/FordY1aver^2], Par1 = AlphaparIN[[1]], *)
(**)
(*   Par2 = AlphaparIN[[2]], Par3 = AlphaparIN[[3]], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, *)
(**)
(*  KEYjoin == *)
(**)
(*   5, {model = *)
(**)
(*    FunctE[omega, AlphaparIN[[1]], Z10, Gamma1p, omega2p, Z20, *)
(**)
(*     Gamma2p, Gamma12p], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, Z10 > 0, Gamma1p > 0, omega2p > 0, Z20 > 0, *)
(**)
(*      Gamma2p > 0, *)
(**)
(*      Gamma12p > 0}, {{Z10, AlphaparIN[[2]]}, {Gamma1p, *)
(**)
(*       AlphaparIN[[3]]}, {omega2p, AlphaparIN[[4]]}, {Z20, *)
(**)
(*       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, *)
(**)
(*       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = AlphaparIN[[1]], Par2 = Z10 /. FITtemp["BestFitParameters"],*)
(**)
(*    Par3 = Gamma1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par7 = Gamma12p /. FITtemp["BestFitParameters"]}, *)
(**)
(*  KEYjoin == *)
(**)
(*   6, {model = *)
(**)
(*    FunctE[omega, AlphaparIN[[1]], Z10, Gamma1p, omega2p, Z20, *)
(**)
(*     Gamma2p, Gamma12p], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, Z10 > 0, Gamma1p > 0, omega2p > 0, Z20 > 0, *)
(**)
(*      Gamma2p > 0, *)
(**)
(*      Gamma12p > 0}, {{Z10, AlphaparIN[[2]]}, {Gamma1p, *)
(**)
(*       AlphaparIN[[3]]}, {omega2p, AlphaparIN[[4]]}, {Z20, *)
(**)
(*       AlphaparIN[[5]]}, {Gamma2p, AlphaparIN[[6]]}, {Gamma12p, *)
(**)
(*       AlphaparIN[[7]]}}, omega, Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = AlphaparIN[[1]], Par2 = Z10 /. FITtemp["BestFitParameters"],*)
(**)
(*    Par3 = Gamma1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}, *)
(**)
(*  KEYjoin == *)
(**)
(*   7, {model = *)
(**)
(*    FunctE[omega, omega1p, Z10, Gamma1p, omega2p, Z20, Gamma2p, 0], *)
(**)
(*   FITtemp = *)
(**)
(*    NonlinearModelFit[*)
(**)
(*     Data1, {model, omega1p > 0, Z10 > 0, Gamma1p > 0, omega2p > 0, *)
(**)
(*      Z20 > 0, *)
(**)
(*      Gamma2p > 0}, {{omega1p, AlphaparIN[[1]]}, {Z10, *)
(**)
(*       AlphaparIN[[2]]}, {Gamma1p, AlphaparIN[[3]]}, {omega2p, *)
(**)
(*       AlphaparIN[[4]]}, {Z20, AlphaparIN[[5]]}, {Gamma2p, *)
(**)
(*       AlphaparIN[[6]]}}, omega, Weights -> 1/FordY1aver^2], *)
(**)
(*   Par1 = omega1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par2 = Z10 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par3 = Gamma1p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par4 = omega2p /. FITtemp["BestFitParameters"], *)
(**)
(*   Par5 = Z20 /. FITtemp["BestFitParameters"], *)
(**)
(*   Par6 = Gamma2p /. FITtemp["BestFitParameters"], Par7 = 0}];*)
(**)
(*OUTData ={{Par1},{Par2},{Par3},{Par4},{Par5},{Par6},{Par7}}; *)
(**)
(*Export["outputMathem.dat", OUTData , "Table"];*)
(**)
(*Quit[] *)
