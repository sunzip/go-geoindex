var init = function(update, onInit) {
    var prmstr = window.location.search.substr(1)
    var prmarr = prmstr.split("&");
    var params = {};

    for ( var i = 0; i < prmarr.length; i++) {
        var tmparr = prmarr[i].split("=");
        params[tmparr[0]] = tmparr[1];
    }
    var refreshArray=[]

    // 拖动时，会调用多次refresh，卡顿
    // 需要做下性能优化
    var refreshCache = function() {
        //清除之前的请求
        for ( var i=0;i<refreshArray.length;i++){
            window.clearTimeout(refreshArray[i])
        }

        r= window.setTimeout(refresh,100) //0.1 s
        refreshArray.push(r)
    }
    var refresh = function() {
        var bounds = map.getBounds();

        var ne = bounds.northeast;
        var sw = bounds.southwest;

        var url = '/pointsShanghai?topLeftLat=' + ne.lat + '&topLeftLon=' + sw.lng + '&bottomRightLat=' + sw.lat + '&bottomRightLon=' + ne.lng + '&index=' + params["index"];

        console.log(url);

        $.getJSON(url,
            function(data) {
                update(data);
             }
        );
    }

    function initialize() {

        var marker
        marker, map = new AMap.Map("map-canvas", {
            resizeEnable: true,
            center: [121.219958,31.096862],
            zoom: 13
        });
        // map = new google.maps.Map(document.getElementById('map-canvas'), mapOptions);
        // google.maps.event.addListener(map, 'idle', refresh);
        // todo:
        
        // map.on('movestart', mapMovestart);
        // map.on('mapmove', refresh);
        map.on('moveend', refreshCache);

        if (onInit) {
            onInit(map, params)
            map.on('complete', function() {
                refresh();
            });
        }
    }

    if (params["refresh"]) {
        var refreshCycle = function() {
            refresh();
            setTimeout(refreshCycle, 1000);
        }
        setTimeout(refreshCycle, 2000);
    }

    $(document).ready(initialize);
}

function mapMovestart(){
    document.querySelector("#text").innerText = '地图移动开始';
}
function mapMove(){
    logMapinfo();
    document.querySelector("#text").innerText = '地图正在移动';
}
function mapMoveend(){
    document.querySelector("#text").innerText = '地图移动结束';
}
var logMapinfo = function (){
    var zoom = map.getZoom(); //获取当前地图级别
    var center = map.getCenter(); //获取当前地图级别
    document.querySelector("#map-center").innerText = center.toString();
};