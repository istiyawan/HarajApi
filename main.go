package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HarajApi struct {
	body         string `json:"body"`
	city         string `json:"city"`
	commentCount int    `json:"commentCount"`
	date         int    `json:"date"`
	thumbURL     string `json:"thumbURL"`
	title        string `json:"title"`
	username     string `json:"username"`
}

var harajAPI = []HarajApi{
	{
		title:        "ددسن 86 غمارتين",
		username:     "ffaris7",
		thumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/900x506_40794ac6-4d70-4b56-95cd-40be11922c44.jpg",
		commentCount: 11,
		city:         "الطايف",
		date:         1621848880,
		body:         "السيارة: نيسان - ددسن\nالموديل: 1986\nبدون لوحات وبدون استمارة\nحالة السيارة: مستعملة\nالقير: قير عادي\nنوع الوقود: بنزين\nالسيارة لا تشكو من شي \nبدون لوحات واستماره \nتنفع للمزارع والحلال",
	},
	{
		title:        "بترول بدون حراره ولا مصروف بنزين",
		username:     "faisal taha",
		thumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/142x288-1_-ZZo3ZDrRMaWuu9.jpg",
		commentCount: 1,
		city:         "الرياض",
		date:         1621848858,
		body:         "-السيارة:  نيسان - باترول      \n-الموديل:  2010                \n-الحالة:  مستعملة\nنقوم بإلغاء المروحه الاساسيه وتركيب مكانها مروحة كهرب جديده في الكرتون اصلي امريكي\nضفيره خاصه للمروحه جودة عالية جدا\nفوائدها\nحراره ثابه مع الزحمات والوقوف الطويل\nبدون صوت ضوجان وحنه\nعزم ناااااار خصوصا القومه\nتوفير بنزين بشكل ملحوض\nوعلى الضمان\nجمل محرك بحث\nتركيب مراوح ترهيم مراوح حراره تعديل مراوح\nمراوح كهربائية تفصيل مراوح كهرب\nالرياض صناعية الخليج\nابو فيصل الخفاجي",
	},
	{
		title:        "النترا 2017 فل كامل",
		username:     "thoo9",
		thumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x599-1_-GO__MTYyMTg0ODY1MzQ4ODcyMDEzMTgw.jpg",
		commentCount: 1,
		city:         "جده",
		date:         1621848855,
		body:         "- النوع : هونداي - النترا\n- الموديل: 2017\n- حالة السيارة : مستعملة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n- الماكينة : كبيرة 2000 CC\n- الممشى :  244 ألف\n- تشغيل بصمة \n- دخول وفتح الباب بصمة \n- مواصفات الفل \nسيارة بحالة ممتازة ونظيفة بودي وكالة\nاشتريتها من سنتين تقريبا للاهل استخدام حشمه سبب البيع الرغبة في سيارة جيب مرتفعة\nفحصتها وقت الشراء ولا يوجد فيها رشوش وغيره حسب الفحص و الشخص اللي اشتريت منه\nالسيارة مجددة ومفحوصة ومأمن عليها قبل شهر تقريبا\nالبيع بسعر 33 ألف\nالتواصل هنا ع العام\nالرجاء فقط الجادين والراغب في الشراء بالسعر المذكور اعلاه\nدمتم سالمين\n            السعر :33000",
	},
	{
		title:        "حفرالباطن العزيزيه",
		username:     "ابو عبد العزيز سعود",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x760-1_-623IE63emdWzeY.jpg",
		commentCount: 12,
		city:         "حفر الباطن",
		date:         1621848846,
		body:         "-السيارة:  نيسان - ددسن        \n-الموديل:  2015                \n-القير:  عادي                \n-الوقود:  بنزين               \n-الحالة:  مستعملة             \n-العداد\t: 224 ألف كم\nشرط   البدي   والمحركات    نظيف   ولايشكو   من  اي  شي  ولله  الحمد",
	},
	{
		title:        "فرن كهربائي",
		username:     "سبحان الله وبحمده 78931",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/722x900_20eaa23d-9923-4a3b-81e5-6430cea9e6e1.jpg",
		commentCount: 0,
		city:         "الرياض",
		date:         1620743395,
		body:         "عندي فرن كهربائي شغال ونظيف بس القزاز مكسور الي عند الباب وله تصليح الفرن ممتاز ويستحمل حراره قويه  للبيع على السوم واتس 0504306174",
	},
	{
		title:        "هونداي كريتا 2021 نص فل ميد",
		username:     "khalid 7566",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x637-1_-GO__MTYyMTIzMzY3Nzk4NzM4ODI5NzUyNA.jpg",
		commentCount: 1,
		city:         "حائل",
		date:         1621234623,
		body:         "- النوع : هونداي - كريتا\n- الموديل: 2021\n- حالة السيارة : وكالة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n الموديل: هونداي - كريتا نص فل مــيد\n- حالة السيارة : وكالة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n-معرض الشـــــــــــرق اللامــــــــــــــــع للسيـــــــــــارات\n            حائــــــــــــــــــــل\n-محرك 1500 سي سي\n-زجاج كهرباء\n-جير أتوماتك\n-دفع أمامي جنوط كروم\n-أشارات بالمرايات - مقاعد مخمل - تشغيل بصمة \n-كاميرا خلفية -فتجة سقف بانروما\nللأستفسار عن السيارة الرجاء التواصل مع قسم المبيعات مع مراعاة وقت الأتصال :-\n0558070751",
	},
	{
		title:        "شاشة سمارت 55بوصة اسكاي ورث كل شي موضح بالكرت",
		username:     "محبه الحرمين 926",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/772x772-1_-NLLxUvYvj5GElQ.jpg",
		commentCount: 7,
		city:         "مكه",
		date:         1621848889,
		body:         "ثلاجة شارب تبريد ممتاز 600 شاشة 1200 السوم وصل 900 ولا بعت",
	},
	{
		title:        "شريط سوني 4",
		username:     "الجهني الرفاعي",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/675x900_dde84f20-5326-46c3-8acb-d4231fb8239a.jpg",
		commentCount: 0,
		city:         "ينبع",
		date:         1621489418,
		body:         "ورحمته الله   :️  شريط سوني:4  أسم الشريط :THE LAST OF US PART II  جديد ✅   السعر : 240             السعر :240",
	},
	{
		title:        "تلفزيون صغير لاصحاب الاثري",
		username:     "الجيل الرقمي",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1479-1_-GO__MTYxOTcwODA0ODI4MjU3MjM1MTkwMg.jpg",
		commentCount: 0,
		city:         "مكه",
		date:         1619708137,
		body:         "تلفزيون صغير ما ادري اذا شغال او لا\nالبيع ع السوم برقم الجوال",
	},
	{
		title:        "جيب لكزس للبيع موديل 2016 فل كامل",
		username:     "ابو تركي123321",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/506x900-1_-7fHoaBNE8fA4kw.jpg",
		commentCount: 7,
		city:         "الرياض",
		date:         1618675451,
		body:         "-السيارة:  لكزس - LX           &lt;br&gt;-الموديل:  2016                &lt;br&gt;-الحالة:  مستعملة&lt;br&gt;يا اخوان انا مجرد معلن والتواصل مع صاحب الجوال . اذكروا الله وصلوا على النبي محمد عليه الصلاة والسلام  الاعلان  جيب لكزس موديل 2016 فل كامل الدفعة الاولى 2016 وصاحبه ذكر لي بأنه يوجد به عفط بسيط في تكاية الاقدام لركوب السياره لايرى إلا بتقريب الصورة وممشاه الحقيقي 128000 كيلوا الان  كما يوجد فديو لمن أراد رؤية الداخلي وهو صامل يشتري كما ارجوا وضع رقم الجوال للمصداقية ولن ينظر للسومة بدون رقم الجوال بالعام يا اخوان الحراج او التواصل واتس آب والله يوفق الجميع لكل خير",
	},
	{
		title:        "Kittens for sale with their mom 3 قطط للبيع",
		username:     "yusuf malik",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/1008x756-1_-UiSPlBdUaHsNg7.jpg",
		commentCount: 0,
		city:         "الرياض",
		date:         1621848888,
		body:         "3 kittens for sale with their mom.\n*Age: The 3 kittens 46 days and the mother is 1 year and 4 months \n* The big cat is vaccinatined \nThe kittens are so playful and they love kids, very healthy and they started to eat dry food! The big cat is so lovely and healthy she eats canes food and fry food. \nThe price of white kittens (females) is 1,500 each.\nThe price of the tiger (female) one is 1,500 \nAnd for the big cat (female) 1,000 \nFor more information please contact me on WhatsApp. Thank you",
	},
	{
		title:        "اي عبوة معروضه فقط بسعر 99 ريال لفتره محدودة",
		username:     "salmafarid",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1400-1_-GO__MTYxNjM5Mzk0Mzk5MjkzMTg4NTkxOA.jpg",
		commentCount: 4,
		city:         "جده",
		date:         1619890507,
		body:         "اذا كنت من عشاق رائحة الفانيليا ..\nفليس هناك اجمل ولا افخم من هذا العطر ..\nثبات وفوحان جدا ممتاز .\nوفوق هذا خفيف ومو مزعج \nعطر مانسيرا روز فانيلا : هو مزيج مابين الروز والفانيلا ، ويمتاز برائحته الجميلة وفوحانة الرائع ، مناسب للاحداث اليومية والرسمية .\nالحجم /100 مل\n            السعر :99",
	},
	{
		title:        "بطريه شحن هواوي",
		username:     "abofahad_342",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x780-1_-fOaX3DGUN6W5Df.jpg",
		commentCount: 0,
		city:         "بريدة",
		date:         1621848887,
		body:         "يوجد بطريه شحن جوال نضيفه وجديده استخدام بسيط بطريه حقت أندرويد هواوي وجلكسي شحن سريع الموقع بريده البيع مستعجل تخزن نسبت 200 ٪السعر 30 ريال\n            السعر :30",
	},
	{
		title:        "الاكاله الذكيه 25 ريال",
		username:     "وائل شعبان محمد",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x480-1_-6l0SGfzZTdC3Uh.jpg",
		commentCount: 0,
		city:         "مكه",
		date:         1621848886,
		body:         "الاكاله الذكيه 25 ريال\n1-تفصل القشر عن الاكل عن طريق درج اسفل الاكاله لمنع اهدار الحب وقت التنظيف \n2-تسع نصف كيلو حب مما يكفى الطائر 9 ايام\n3-يوجد بها علاقه داخليه وخارجيه مما يتيح لك تركيبها داخل القفص او خارجه\nيوجد توصيل لخارج مكه فقط\nعن طريق شركه زاجل للشحن \nداخل مكه راسلنى واتس ارسلك اللكيشن \nللطلب والاستفسار تواصل واتس على&nbsp; \n0536142062\n            السعر :25",
	},
	{
		title:        "حساب قراند شبه رباعي للبيع",
		username:     "0500078226",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1479-1_-GO__MTYyMTg0ODQ4OTM5NzU3MzM5NjUwMw.jpg",
		commentCount: 0,
		city:         "الطايف",
		date:         1621848884,
		body:         ":اسلام عليكم ورحمة الله وبركاته \nعندي حساب شبه رباعي للبيع قراند شاري كل شي بلعبه وفيه دلكسوتنباع الوحده بي اثنين مليون  ثمن ميت الف تقدر تبيع باليوم ثلاثه وفي دباب يطر ينباع الواحد بي مليون وميتن الف وفيه ملا بس مهكره وفيه شخصيه الولد نفس كل شي في شخصيه البنت",
	},
	{
		"title":        "للبيع 6 معيز مع تيس",
		"username":     "mr.saodi",
		"thumbURL":     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/674x900_aa42f986-e0ee-4ca5-9e02-4fa01d84237f.jpg",
		"commentCount": 0,
		"city":         "الدلم",
		"date":         1621848883,
		"body":         "للبيع 6معيز ماهن دفيع مع تيس السن تام  الحلال ضعيف لكن شرط الصحة  الموقع الدلم  السوم برقم الجوال لاهنتم",
	},
	{
		title:        "إدارة المتاجر الالكترونية",
		username:     "سهولة الوصول",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x800-1_-GO__MTYxOTc3MDE5OTcxODIzNzI0MTkwNg.jpg",
		commentCount: 0,
		city:         "الرياض",
		date:         1620028787,
		body:         "طوّر متجرك الإلكتروني واجعله مصدر دخلك الأساسي\n1. دراسة السوق وتحليل المنافسين\n2. تحديد الفئات المستهدفة\n3. تحديد منصات التواصل الاجتماعي الأكثر ملائمة للنشاط التجاري\n4. اختيار العروض الأكثر ملائمة للجمهور المستهدف والايام والفعاليات السنوية\n5. إنشاء خطة استراتيجية متكاملة لإدارة المتجر الإلكتروني\n6. كتابة المحتوى المناسب لصفحات الهبوط الخاصة بالعروض\n7. الدعم الفني\nتواصل معنا عبر الواتس اب:\n0558003787",
	},
	{
		title:        "ساعات أطقم رجالي ونسائي",
		username:     "عالم النحلة",
		thumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/480x640-1_-J5O3opoUD8LCsJ.jpg",
		commentCount: 0,
		city:         "جده",
		date:         1619655809,
		body:         "🌹عروض العيد 🌹\nخصم يصل إلى 50%\n#أطقم ساعات رجالية أنيقة كشخة الساعات الأكثر رغبة وطلبا🎁\n#ساعة ⌚\n# قلم 🖋️\n#سبحة\n#كبك\n#علبة\n#كرت ضمان\nسعر الطقم 100 غير شامل التوصيل\n*للعلم *ضمان لمدة سنة من تاريخ الشراء\n_____________________________________\n#اطقم ساعات نسائية أنيقة كشخة الساعات الأكثر رغبة وطلبا... \n#ساعة \n#ثلاث اسورات \n#علبة \n#كرت \n# ضمان \nسعر الطقم 100ريال فقط. \n*للعلم *ضمان لمدة سنة من تاريخ الشراء \n_____________________________________\nللطلب والاستفسار التواصل على الرقم 0577802537 📲",
	},
}

func getHarajCar(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, harajAPI)
}

func getHarajCarByTitle(c *gin.Context) {
	title := c.Param("title")
	fmt.Println("title", title)

	for _, car := range harajAPI {
		if car.title == title {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Car Not Found"})

}

func main() {
	router := gin.Default()
	router.GET("/", getHarajCar)
	router.GET("/:title", getHarajCarByTitle)
	// router.Run("localhost:8080")
	router.Run(":" + os.Getenv("PORT"))
}
