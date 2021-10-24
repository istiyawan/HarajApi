package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HarajApi struct {
	Body         string `json:"body"`
	City         string `json:"city"`
	CommentCount int    `json:"commentCount"`
	Date         int    `json:"date"`
	ThumbURL     string `json:"thumbURL"`
	Title        string `json:"title"`
	Username     string `json:"username"`
}

var harajAPI = []HarajApi{
	{
		Title:        "ددسن 86 غمارتين",
		Username:     "ffaris7",
		ThumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/900x506_40794ac6-4d70-4b56-95cd-40be11922c44.jpg",
		CommentCount: 11,
		City:         "الطايف",
		Date:         1621848880,
		Body:         "السيارة: نيسان - ددسن\nالموديل: 1986\nبدون لوحات وبدون استمارة\nحالة السيارة: مستعملة\nالقير: قير عادي\nنوع الوقود: بنزين\nالسيارة لا تشكو من شي \nبدون لوحات واستماره \nتنفع للمزارع والحلال",
	},
	{
		Title:        "بترول بدون حراره ولا مصروف بنزين",
		Username:     "faisal taha",
		ThumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/142x288-1_-ZZo3ZDrRMaWuu9.jpg",
		CommentCount: 1,
		City:         "الرياض",
		Date:         1621848858,
		Body:         "-السيارة:  نيسان - باترول      \n-الموديل:  2010                \n-الحالة:  مستعملة\nنقوم بإلغاء المروحه الاساسيه وتركيب مكانها مروحة كهرب جديده في الكرتون اصلي امريكي\nضفيره خاصه للمروحه جودة عالية جدا\nفوائدها\nحراره ثابه مع الزحمات والوقوف الطويل\nبدون صوت ضوجان وحنه\nعزم ناااااار خصوصا القومه\nتوفير بنزين بشكل ملحوض\nوعلى الضمان\nجمل محرك بحث\nتركيب مراوح ترهيم مراوح حراره تعديل مراوح\nمراوح كهربائية تفصيل مراوح كهرب\nالرياض صناعية الخليج\nابو فيصل الخفاجي",
	},
	{
		Title:        "النترا 2017 فل كامل",
		Username:     "thoo9",
		ThumbURL:     "https://s4-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x599-1_-GO__MTYyMTg0ODY1MzQ4ODcyMDEzMTgw.jpg",
		CommentCount: 1,
		City:         "جده",
		Date:         1621848855,
		Body:         "- النوع : هونداي - النترا\n- الموديل: 2017\n- حالة السيارة : مستعملة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n- الماكينة : كبيرة 2000 CC\n- الممشى :  244 ألف\n- تشغيل بصمة \n- دخول وفتح الباب بصمة \n- مواصفات الفل \nسيارة بحالة ممتازة ونظيفة بودي وكالة\nاشتريتها من سنتين تقريبا للاهل استخدام حشمه سبب البيع الرغبة في سيارة جيب مرتفعة\nفحصتها وقت الشراء ولا يوجد فيها رشوش وغيره حسب الفحص و الشخص اللي اشتريت منه\nالسيارة مجددة ومفحوصة ومأمن عليها قبل شهر تقريبا\nالبيع بسعر 33 ألف\nالتواصل هنا ع العام\nالرجاء فقط الجادين والراغب في الشراء بالسعر المذكور اعلاه\nدمتم سالمين\n            السعر :33000",
	},
	{
		Title:        "حفرالباطن العزيزيه",
		Username:     "ابو عبد العزيز سعود",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x760-1_-623IE63emdWzeY.jpg",
		CommentCount: 12,
		City:         "حفر الباطن",
		Date:         1621848846,
		Body:         "-السيارة:  نيسان - ددسن        \n-الموديل:  2015                \n-القير:  عادي                \n-الوقود:  بنزين               \n-الحالة:  مستعملة             \n-العداد\t: 224 ألف كم\nشرط   البدي   والمحركات    نظيف   ولايشكو   من  اي  شي  ولله  الحمد",
	},
	{
		Title:        "فرن كهربائي",
		Username:     "سبحان الله وبحمده 78931",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/722x900_20eaa23d-9923-4a3b-81e5-6430cea9e6e1.jpg",
		CommentCount: 0,
		City:         "الرياض",
		Date:         1620743395,
		Body:         "عندي فرن كهربائي شغال ونظيف بس القزاز مكسور الي عند الباب وله تصليح الفرن ممتاز ويستحمل حراره قويه  للبيع على السوم واتس 0504306174",
	},
	{
		Title:        "هونداي كريتا 2021 نص فل ميد",
		Username:     "khalid 7566",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x637-1_-GO__MTYyMTIzMzY3Nzk4NzM4ODI5NzUyNA.jpg",
		CommentCount: 1,
		City:         "حائل",
		Date:         1621234623,
		Body:         "- النوع : هونداي - كريتا\n- الموديل: 2021\n- حالة السيارة : وكالة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n الموديل: هونداي - كريتا نص فل مــيد\n- حالة السيارة : وكالة\n- القير:  اوتوماتيك\n- الوقود : بنزين\n-معرض الشـــــــــــرق اللامــــــــــــــــع للسيـــــــــــارات\n            حائــــــــــــــــــــل\n-محرك 1500 سي سي\n-زجاج كهرباء\n-جير أتوماتك\n-دفع أمامي جنوط كروم\n-أشارات بالمرايات - مقاعد مخمل - تشغيل بصمة \n-كاميرا خلفية -فتجة سقف بانروما\nللأستفسار عن السيارة الرجاء التواصل مع قسم المبيعات مع مراعاة وقت الأتصال :-\n0558070751",
	},
	{
		Title:        "شاشة سمارت 55بوصة اسكاي ورث كل شي موضح بالكرت",
		Username:     "محبه الحرمين 926",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/772x772-1_-NLLxUvYvj5GElQ.jpg",
		CommentCount: 7,
		City:         "مكه",
		Date:         1621848889,
		Body:         "ثلاجة شارب تبريد ممتاز 600 شاشة 1200 السوم وصل 900 ولا بعت",
	},
	{
		Title:        "شريط سوني 4",
		Username:     "الجهني الرفاعي",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/675x900_dde84f20-5326-46c3-8acb-d4231fb8239a.jpg",
		CommentCount: 0,
		City:         "ينبع",
		Date:         1621489418,
		Body:         "ورحمته الله   :️  شريط سوني:4  أسم الشريط :THE LAST OF US PART II  جديد ✅   السعر : 240             السعر :240",
	},
	{
		Title:        "تلفزيون صغير لاصحاب الاثري",
		Username:     "الجيل الرقمي",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1479-1_-GO__MTYxOTcwODA0ODI4MjU3MjM1MTkwMg.jpg",
		CommentCount: 0,
		City:         "مكه",
		Date:         1619708137,
		Body:         "تلفزيون صغير ما ادري اذا شغال او لا\nالبيع ع السوم برقم الجوال",
	},
	{
		Title:        "جيب لكزس للبيع موديل 2016 فل كامل",
		Username:     "ابو تركي123321",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/506x900-1_-7fHoaBNE8fA4kw.jpg",
		CommentCount: 7,
		City:         "الرياض",
		Date:         1618675451,
		Body:         "-السيارة:  لكزس - LX           &lt;br&gt;-الموديل:  2016                &lt;br&gt;-الحالة:  مستعملة&lt;br&gt;يا اخوان انا مجرد معلن والتواصل مع صاحب الجوال . اذكروا الله وصلوا على النبي محمد عليه الصلاة والسلام  الاعلان  جيب لكزس موديل 2016 فل كامل الدفعة الاولى 2016 وصاحبه ذكر لي بأنه يوجد به عفط بسيط في تكاية الاقدام لركوب السياره لايرى إلا بتقريب الصورة وممشاه الحقيقي 128000 كيلوا الان  كما يوجد فديو لمن أراد رؤية الداخلي وهو صامل يشتري كما ارجوا وضع رقم الجوال للمصداقية ولن ينظر للسومة بدون رقم الجوال بالعام يا اخوان الحراج او التواصل واتس آب والله يوفق الجميع لكل خير",
	},
	{
		Title:        "Kittens for sale with their mom 3 قطط للبيع",
		Username:     "yusuf malik",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/1008x756-1_-UiSPlBdUaHsNg7.jpg",
		CommentCount: 0,
		City:         "الرياض",
		Date:         1621848888,
		Body:         "3 kittens for sale with their mom.\n*Age: The 3 kittens 46 days and the mother is 1 year and 4 months \n* The big cat is vaccinatined \nThe kittens are so playful and they love kids, very healthy and they started to eat dry food! The big cat is so lovely and healthy she eats canes food and fry food. \nThe price of white kittens (females) is 1,500 each.\nThe price of the tiger (female) one is 1,500 \nAnd for the big cat (female) 1,000 \nFor more information please contact me on WhatsApp. Thank you",
	},
	{
		Title:        "اي عبوة معروضه فقط بسعر 99 ريال لفتره محدودة",
		Username:     "salmafarid",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1400-1_-GO__MTYxNjM5Mzk0Mzk5MjkzMTg4NTkxOA.jpg",
		CommentCount: 4,
		City:         "جده",
		Date:         1619890507,
		Body:         "اذا كنت من عشاق رائحة الفانيليا ..\nفليس هناك اجمل ولا افخم من هذا العطر ..\nثبات وفوحان جدا ممتاز .\nوفوق هذا خفيف ومو مزعج \nعطر مانسيرا روز فانيلا : هو مزيج مابين الروز والفانيلا ، ويمتاز برائحته الجميلة وفوحانة الرائع ، مناسب للاحداث اليومية والرسمية .\nالحجم /100 مل\n            السعر :99",
	},
	{
		Title:        "بطريه شحن هواوي",
		Username:     "abofahad_342",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x780-1_-fOaX3DGUN6W5Df.jpg",
		CommentCount: 0,
		City:         "بريدة",
		Date:         1621848887,
		Body:         "يوجد بطريه شحن جوال نضيفه وجديده استخدام بسيط بطريه حقت أندرويد هواوي وجلكسي شحن سريع الموقع بريده البيع مستعجل تخزن نسبت 200 ٪السعر 30 ريال\n            السعر :30",
	},
	{
		Title:        "الاكاله الذكيه 25 ريال",
		Username:     "وائل شعبان محمد",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/360x480-1_-6l0SGfzZTdC3Uh.jpg",
		CommentCount: 0,
		City:         "مكه",
		Date:         1621848886,
		Body:         "الاكاله الذكيه 25 ريال\n1-تفصل القشر عن الاكل عن طريق درج اسفل الاكاله لمنع اهدار الحب وقت التنظيف \n2-تسع نصف كيلو حب مما يكفى الطائر 9 ايام\n3-يوجد بها علاقه داخليه وخارجيه مما يتيح لك تركيبها داخل القفص او خارجه\nيوجد توصيل لخارج مكه فقط\nعن طريق شركه زاجل للشحن \nداخل مكه راسلنى واتس ارسلك اللكيشن \nللطلب والاستفسار تواصل واتس على&nbsp; \n0536142062\n            السعر :25",
	},
	{
		Title:        "حساب قراند شبه رباعي للبيع",
		Username:     "0500078226",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x1479-1_-GO__MTYyMTg0ODQ4OTM5NzU3MzM5NjUwMw.jpg",
		CommentCount: 0,
		City:         "الطايف",
		Date:         1621848884,
		Body:         ":اسلام عليكم ورحمة الله وبركاته \nعندي حساب شبه رباعي للبيع قراند شاري كل شي بلعبه وفيه دلكسوتنباع الوحده بي اثنين مليون  ثمن ميت الف تقدر تبيع باليوم ثلاثه وفي دباب يطر ينباع الواحد بي مليون وميتن الف وفيه ملا بس مهكره وفيه شخصيه الولد نفس كل شي في شخصيه البنت",
	},
	{
		Title:        "للبيع 6 معيز مع تيس",
		Username:     "mr.saodi",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/674x900_aa42f986-e0ee-4ca5-9e02-4fa01d84237f.jpg",
		CommentCount: 0,
		City:         "الدلم",
		Date:         1621848883,
		Body:         "للبيع 6معيز ماهن دفيع مع تيس السن تام  الحلال ضعيف لكن شرط الصحة  الموقع الدلم  السوم برقم الجوال لاهنتم",
	},
	{
		Title:        "إدارة المتاجر الالكترونية",
		Username:     "سهولة الوصول",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/800x800-1_-GO__MTYxOTc3MDE5OTcxODIzNzI0MTkwNg.jpg",
		CommentCount: 0,
		City:         "الرياض",
		Date:         1620028787,
		Body:         "طوّر متجرك الإلكتروني واجعله مصدر دخلك الأساسي\n1. دراسة السوق وتحليل المنافسين\n2. تحديد الفئات المستهدفة\n3. تحديد منصات التواصل الاجتماعي الأكثر ملائمة للنشاط التجاري\n4. اختيار العروض الأكثر ملائمة للجمهور المستهدف والايام والفعاليات السنوية\n5. إنشاء خطة استراتيجية متكاملة لإدارة المتجر الإلكتروني\n6. كتابة المحتوى المناسب لصفحات الهبوط الخاصة بالعروض\n7. الدعم الفني\nتواصل معنا عبر الواتس اب:\n0558003787",
	},
	{
		Title:        "ساعات أطقم رجالي ونسائي",
		Username:     "عالم النحلة",
		ThumbURL:     "https://s3-eu-west-1.amazonaws.com/img4.haraj.com.sa/cache4/480x640-1_-J5O3opoUD8LCsJ.jpg",
		CommentCount: 0,
		City:         "جده",
		Date:         1619655809,
		Body:         "🌹عروض العيد 🌹\nخصم يصل إلى 50%\n#أطقم ساعات رجالية أنيقة كشخة الساعات الأكثر رغبة وطلبا🎁\n#ساعة ⌚\n# قلم 🖋️\n#سبحة\n#كبك\n#علبة\n#كرت ضمان\nسعر الطقم 100 غير شامل التوصيل\n*للعلم *ضمان لمدة سنة من تاريخ الشراء\n_____________________________________\n#اطقم ساعات نسائية أنيقة كشخة الساعات الأكثر رغبة وطلبا... \n#ساعة \n#ثلاث اسورات \n#علبة \n#كرت \n# ضمان \nسعر الطقم 100ريال فقط. \n*للعلم *ضمان لمدة سنة من تاريخ الشراء \n_____________________________________\nللطلب والاستفسار التواصل على الرقم 0577802537 📲",
	},
}

func getHarajCar(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, harajAPI)
}

func getHarajCarByTitle(c *gin.Context) {
	title := c.Param("title")
	fmt.Println("title", title)

	for _, car := range harajAPI {
		if car.Title == title {
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
