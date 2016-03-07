metropolis version 2.0

2016/3/5

1. 開発・実行環境

  [CPU] Intel(R) Core(TM) i7-5960X CPU/3.0GHz
   [OS] Linux version 4.4.0/64 bit 
  [Distribution] Ubuntu 16.04
  [golang] go version 1.6

2. プログラムの使用方法

2.1 プログラム名

  metropolis_V2

2.2 プログラムのオプションスイッチとデフォルト値

  $ ./metropolis_V2 -h

  Usage of metropolis_V2:

    ./metropolis_V2 [options] -N Step_Number

    -Cout string
        Cout file. (default "Cout.csv")
    -DataDir string
        Input data directory. (default "./data")
    -Eout string
        Eout file. (default "Eout.dat")
    -N float
        Step number.
    -Temp string
        For parallel tempering. (default "200,300,10")

2.3 実行例

  [ステップ数] 100000(1e5)
        [温度] 100 K から  35 K 刻みで 500K まで

    $ ./metropolis_V2 -N 100000 -T100,500,35 2>&1 | tee log

  ステップ数には指数表記を使う事も可能。

    $ ./metropolis_V2 -N 1e5 -T100,500,35 2>&1 | tee log

3. 入出力データ

3.1 入力データ

  入力データは "-DataDir" オプションで指定するディレクトリに配置する。
  (デフォルトでは "./data" ディレクトリ)

  ./data
    ├── CcoordsAVE.csv
    ├── HcoordsAVE.csv
    ├── BrcoordsAVE.csv
    ├── PrecursorUnitCellAxes.csv
    ├── UnitCell2.csv
    ├── Lattice.csv
    ├── Character.csv
    ├── AdjCuml.csv
    ├── KernelregSAtt.json
    ├── KernelregSRepLog.json
    ├── SvmModelNzp.json
    ├── SvmModelOp.json
    ├── SvmModelUsp.json
    └── Xeigpc.json

- CcoordsAVE.csv
  HcoordsAVE.csv
  BrcoordsAVE.csv

  Coordinates of atoms in molecule.

  これらのファイルは事前に用意される。metropolis_V2 プログラムの実行前に
  以下の様にシンボリックリンクファイルを作成しておく。

    lrwxrwxrwx 1 Molecule_01.csv -> CcoordsAVE.csv
    lrwxrwxrwx 1 Molecule_02.csv -> HcoordsAVE.csv
    lrwxrwxrwx 1 Molecule_03.csv -> BrcoordsAVE.csv

  シンボリックリンクのファイル名は Molecule_0[1-3].csv とする。

    $ cd data
    $ ln -s CcoordsAVE.csv Molecule_01.csv
    $ ln -s HcoordsAVE.csv Molecule_02.csv
    $ ln -s BrcoordsAVE.csv Molecule_03.csv

  CH3 の様な 2 分子の場合、シンボリックリンクファイルは 2 個になる。

    $ cd data
    $ ln -s CcoordsAVE.csv Molecule_01.csv
    $ ln -s HcoordsAVE.csv Molecule_02.csv

- PrecursorUnitCellAxes.csv
  
  Lattice Vectors.

  事前に用意される。

- UnitCell2.csv

  Unit cell.

  R 上で以下を実行して生成する。

    write.table(format(UnitCell2, digits=22, trim=T), file="UnitCell2.csv",
                sep=",", row.names=FALSE, col.names=FALSE, quote=F)

- Lattice.csv

  Lattice based on the unit cell.

  R 上で以下を実行して生成する。

    write.table(format(Lattice, digits=22, trim=T),
                file="Lattice.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)

- Character.csv

  Specify which unit cell point the lattice point corresponds to.

  R 上で以下を実行して生成する。

    write.table(t(as.matrix(Character-1)),
                file="Character.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)

- AdjCuml.csv

  Sequence of power matrices. Adjacency matrix for the unit cells.

  R 上で以下を実行して生成する。

    writeListData(AdjCuml, "AdjCuml.csv")

- KernelregSAtt.json
  KernelregSRepLog.json

  Kernel-based regularised least-squares regression objects.

  R 上で以下を実行して生成する。

  write(toJSON(unclass(kernelregS_Rep_log),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE), "KernelregSRepLog.json")
  write(toJSON(unclass(kernelregS_Att),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE), "KernelregSAtt.json")

- SvmModelNzp.json
  SvmModelOp.json
  SvmModelUsp.json

  SVM(Support Vector Machines) objects.

  R 上で以下を実行して生成する。

  svm_modelOP$rho <- I(svm_modelOP$rho)
  write(toJSON(unclass(svm_modelOP),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE, force=TRUE), "SvmModelOp.json")
  svm_modelNZP$rho <- I(svm_modelNZP$rho)
  write(toJSON(unclass(svm_modelNZP),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE, force=TRUE), "SvmModelNzp.json")
  svm_modelUSP$rho <- I(svm_modelUSP$rho)
  write(toJSON(unclass(svm_modelUSP),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE, force=TRUE), "SvmModelUsp.json")

- Xeigpc.json

  Prcomp objects.

  R 上で以下を実行して生成する。

  write(toJSON(unclass(XeigPC),
        auto_unbox=TRUE, digits=I(22), pretty=TRUE), "Xeigpc.json")
  
3.2 出力データ

- Cout.csv

  Canonical data.

  metropolis_V2 プログラムの -Cout オプションで変更可能。

  以下を実行して R に読み込む。

    Cout <- loadCoutV2("Cout.csv", Nparallel)

- Eout.dat

  Energy data.

  metropolis_V2 プログラムの -Eout オプションで変更可能。

  以下を実行して R に読み込む。

    Eout <- loadEout("Eout.dat", Nparallel)

- ログファイル

  metropolis_V2 プログラムは処理の経過を標準エラー出力(stderr)に出力する。

  [出力例]
  2016/02/19 14:50:46 N = 300000
  2016/02/19 14:50:46 TempS = []float64{200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300}
  2016/02/19 14:50:46 Nparallel = 11
  2016/02/19 14:50:46 Start.
  2016/02/19 14:50:55 n =  100/N = 300000
                       :

  2016/02/21 06:22:06 n = 300000/N = 300000
  2016/02/21 06:22:08 End.
  2016/02/21 06:22:08 Execution time = 142282.06 s/2371.37 m/39.52 h

4. 実行形式プログラムの作成

  metropolis_V2 プログラムのソースコードを展開したディレクトリに移動して
  以下を実行する。

    $ cd src_dir
    $ go build -o metropolis_V2 .

  正常終了すると実行形式ファイル(metropolis_V2)が生成される。また、ソース
  コードが置いてあるディレクトリに Makefile を用意してあるので、以下を実行
  しても可。

    $ cd src_dir
    $ make b
