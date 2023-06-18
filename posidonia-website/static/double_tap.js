function DetectDoubleTap(element){
    function detectDoubleTapClosure() {
    let lastTap = 0;
    let timeout;
    return function detectDoubleTap(event) {
        const curTime = new Date().getTime();
        const tapLen = curTime - lastTap;
        if (tapLen < 500 && tapLen > 0) {
            window.accept_sending = true;
            event.preventDefault();
        } else {
        timeout = setTimeout(() => {
            clearTimeout(timeout);
        }, 500);
        }
        lastTap = curTime;
    };
    }

    /* Regex test to determine if user is on mobile */
    if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
        element.addEventListener('touchend', detectDoubleTapClosure(), { passive: false });
    }
}
