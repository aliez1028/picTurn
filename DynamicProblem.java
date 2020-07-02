package club.buyclass.relationserver;

public class DynamicProblem {

    //背包重量16
    private static int heavy = 16;  //代表背包可装的最大重点
    private static int count = 6;      //代表有多少件物品

    private static int[] ws = {0,2,4,6,4,3,5};  //物品重量
    private static int[] vs = {0,11,5,15,5,8,22};  //物品价格
    private static Integer[][] results = new Integer[count+1][heavy+1]; //用一个二维数组来表示装了前几件物品 ，背包剩余多少重量时 ，背包的最大价值


    public static void main(String[] args) {
        System.out.println(ks());
    }

    private static int ks(){
        for(int i = 0;i< count+1;i++){ //前几件物品
            for (int j = 0;j< heavy +1;j++ ){ //背包剩余重量
                // 当背包剩余重量为0时肯定装不下东西，所以价值为0，当第0件物品也就是没物品，价值也为0
                if (i == 0 || j == 0) {
                    results[i][j] = 0;
                    continue;
                }
                //如果第i个物品的重量大于背包剩余重量，那么肯定装不下，
                //  所以这时候背包可装的最大价值和装了 前i-1个物品 且 背包剩余重量为 j 时，最大价值相同
                if (ws[i] > j) results[i][j] = results[i-1][j];
                else {
                    //如果背包装的下，那么比较一下，
                    // 如果装入这个东西的时候的价值和不装这个东西时的价值，哪个更大
                    int tempValue1 = results[i-1][j-ws[i]]+vs[i];  //装入的话，则当前背包的最大价值为 背包装入前i-1个物品，剩余可装重量减去当前物品重量时的 最大价值 加上 当前 物品的价值
                    int tempValue2 = results[i-1][j];//不装入的话，那么此时背包里物品的价值就是和 装入(i-1个商品时、且背包剩余空间为j时 一样
                    results[i][j] = tempValue1 > tempValue2 ? tempValue1: tempValue2;
                }
            }
        }
        return results[count][heavy];
    }

}
