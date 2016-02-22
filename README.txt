IslandYosou Version 2.0

2016/3/1

1. 開発・実行環境

  [CPU] Intel(R) Core(TM) i7-5960X CPU/3.0GHz
   [OS] Linux version 4.4.0
  [Distribution] Ubuntu 16.04
  [golang] go version devel +5609a48(version 1.6 相当)

2. プログラムの使用方法

* プログラム名

  metropolis_V2

* プログラムのオプションスイッチとデフォルト値

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

* 実行例

  [ステップ数] 100000(1e5)
        [温度] 100 K から  35 K 刻みで 500K まで

  $ ./metropolis_V2 -N 100000 -T100,500,35 | tee log 2>&1

ステップ数には指数表記を使う事も可能。

  $ ./metropolis_V2 -N 1e5 -T100,500,35 | tee log 2>&1

3. 入出力データ

* 入力データ

  入力データは "-DataDir" オプションで指定するディレクトリに配置する。
  (デフォルトでは "./data" ディレクトリ)

  ./data
    ├── AdjCuml.csv
    ├── BrcoordsAVE.csv
    ├── CcoordsAVE.csv
    ├── Character.csv
    ├── HcoordsAVE.csv
    ├── KernelregSAtt.json
    ├── KernelregSRepLog.json
    ├── Lattice.csv
    ├── PrecursorUnitCellAxes.csv
    ├── SvmModel.json
    ├── SvmModelNzp.json
    ├── SvmModelOp.json
    ├── SvmModelUsp.json
    ├── UnitCell.csv
    ├── UnitCell2.csv
    └── Xeigpc.json
    
* 出力データ
