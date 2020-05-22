package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var urls = []string{
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A61E_20170501110842_DeW95N.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A60F_20170501113055_o27Ri5.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501113356_uxEQLJ.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4EB_20170501113814_UZiHfz.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501124035_YE0W8R.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A8D1_20170501135755_PdziZB.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A4CE_20170501134339_DvpX7b.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A469_20170501145045_habTWf.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501123852_VQoTjc.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9B8_20170501125501_z2nR8n.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9C1_20170501145405_CtnAx4.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A913_20170501155546_IUSHMr.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173022_RPjzni.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A2AB_20170501173355_qi035E.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501173711_3i5zYC.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A67B_20170501165200_2lfyc8.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A888_20170501192711_jNVDWi.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A9CE_20170501195005_WJ08YD.mp3",
	"https://qnrecord.ktvsky.com/wow-thunder/record/201705/00E07E00A656_20170504203953_BnriL7.mp3",
}

const PATH = "/home/apple/go/src/I-love-reading/tests/qiniu/download"

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func downloadSong(i int, u string) {
	out, err := os.Create(fmt.Sprintf("%s/%d.mp3", PATH, i))
	handleErr(err)
	defer out.Close()
	resp, err := http.Get(u)
	handleErr(err)
	defer resp.Body.Close()
	n, err := io.Copy(out, resp.Body)
	handleErr(err)
	fmt.Println(n)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	urls := append(urls, urls...)
	fmt.Println(len(urls))
	t1 := time.Now()
	for i, u := range urls {
		wg.Add(1)
		go downloadSong(i, u)
	}
	wg.Wait()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
