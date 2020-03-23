package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	secIDList := []int{1658402, 1372844, 1327828, 1297614, 1308746, 1320025, 1335203, 1689713, 1391773, 1490744, 1779206, 1648238, 1533666, 1146053, 1692379, 1530859, 1725160, 1698705, 1697062, 1486606, 1576486, 1724290, 1419642, 1684492, 1503760, 1762890, 1593648, 1461944, 1357814, 1311455, 1558988, 1503194, 1767951, 1525506, 1475906, 1432862, 1399576, 1339974, 1555679, 1392342, 1688628, 1769067, 1549879, 1455707, 1467838, 1468611, 1645176, 1694353, 1673351, 1340200, 1785610, 1503714, 1398044, 1370830, 1312459, 1461846, 1282260, 1655551, 1004744, 1635481, 1384378, 1677889, 1491984, 1406188, 795051, 1285444, 1276918, 1436331, 1394096, 1780488, 1651762, 1621274, 1542636, 1491965, 1776290, 1762977, 1784485, 1786522, 1789379, 1785082, 1642692, 1747193, 1661182, 1791627, 1759874, 1738680, 1705187, 1741449, 1748374, 1726965, 1763020, 1456014, 1319612, 1307617, 1740656, 1476280, 1177951, 1303377, 1350533, 1371099, 1691752, 1292599, 1308417, 1360988, 1321166, 1346445, 1382391, 1333079, 1448373, 1605812, 1696682, 1477843, 1680492, 1570860, 1656237, 1656235, 1656236, 1519790, 1269832, 1085788, 1315506, 1636423, 1585036, 1502669, 1696147, 1678314, 1517707, 1429007, 1726671, 1532608, 1548582, 1570583, 1332957, 1406344, 1294836, 1734886, 1419940, 1350262, 1419867, 1761486, 1419097, 1137502, 1560783, 1503735, 1301864, 1166129, 1692457, 1570530, 1529264, 1713556, 1705943, 1628181, 1652573, 1657745, 1603043, 1596791, 1554639, 1601179, 1789706, 1703711, 1546256, 1546255, 1546346, 1765920, 1169193, 1623513, 1349516, 1473672, 1775553, 1606096, 1165360, 1418398, 1640698, 1218389, 1165769, 1542702, 1745952, 1718242, 1745953, 1620544, 1400516, 1619171, 1430041, 1652599, 1725773, 1571109, 1410452, 1605232, 1429536, 1330624, 1370732, 1506207, 1349490, 1595562, 1591260, 1605907, 1740645, 1678722, 1678723, 1711630, 1678721, 1663417, 1629934, 1443871}
	var EINList []string

	c := colly.NewCollector()

	c.OnRequest(func(rq *colly.Request) {
		log.Println("Visiting : ", rq.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		EINList = append(EINList, "error")
	})

	c.OnHTML("body > div > div > div:nth-child(5) > div.panel-body > table > tbody > tr:nth-child(1) > td:nth-child(2)", func(e *colly.HTMLElement) {
		data := e.Text
		fmt.Println(data)
		if strings.Contains(data, "EIN") {
			EINList = append(EINList, data)
		} else {
			EINList = append(EINList, "Please do it manually")
		}
	})

	for _, id := range secIDList {
		c.Visit(string("https://sec.report/CIK/" + strconv.Itoa(id)))
	}

	for _, ein := range EINList {
		fmt.Println(ein)
	}
}