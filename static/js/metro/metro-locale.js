(function($){
    $.Metro.currentLocale = 'en';

    if (METRO_LOCALE != undefined)  $.Metro.currentLocale = METRO_LOCALE; else $.Metro.currentLocale = 'en';
    //console.log(METRO_LOCALE, $.Metro.currentLocale);

    $.Metro.Locale = {
        'en': {
            months: [
                "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
                "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
            ],
            days: [
                "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
                "Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"
            ],
            buttons: [
                "Today", "Clear"
            ]
        },
        'fr': {
            months: [
                "Janvier", "Février", "Mars", "Avril", "Peut", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
                "Jan", "Fév", "Mar", "Avr", "Peu", "Jun", "Jul", "Aoû", "Sep", "Oct", "Nov", "Déc"
            ],
            days: [
                "Sunday", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi",
                "Sn", "Ln", "Md", "Mc", "Ju", "Vn", "Sm"
            ],
            buttons: [
                "Aujourd", "Effacer"
            ]
        },
        'ua': {
            months: [
                "Січень", "Лютий", "Березень", "Квітень", "Травень", "Червень", "Липень", "Серпень", "Вересень", "Жовтень", "Листопад", "Грудень",
                "Січ", "Лют", "Бер", "Кві", "Тра", "Чер", "Лип", "Сер", "Вер", "Жов", "Лис", "Гру"
            ],
            days: [
                "Неділя", "Понеділок", "Вівторок", "Середа", "Четвер", "П’ятниця", "Субота",
                "Нд", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"
            ],
            buttons: [
                "Сьогодні", "Очистити"
            ]
        },
        'ru': {
            months: [
                "Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
                "Янв", "Фев", "Мар", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"
            ],
            days: [
                "Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота",
                "Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"
            ],
            buttons: [
                "Сегодня", "Очистить"
            ]
        }
    };

    $.Metro.setLocale = function(locale, data){
        $.Metro.Locale[locale] = data;
    };
})(jQuery);