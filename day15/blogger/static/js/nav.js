let obj;
const As = document.getElementById("topnav").getElementsByTagName('a');
obj = As[0];
for (let i = 1; i < As.length; i++) {
    if (window.location.href.indexOf(As[i].href) >= 0) {
        obj = As[i];
    }
}
obj.id = 'topnav_current'