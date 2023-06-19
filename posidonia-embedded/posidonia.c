#include <curl/curl.h>
#include <stdio.h>
#include <stdlib.h>

int posidonia(void)
{
    CURL *curl = curl_easy_init();

    if(!curl)
    {
        fprintf(stderr, "init failed!\n");
        return EXIT_FAILURE;
    }

    curl_easy_setopt(curl, CURLOPT_URL, "http://youtube.com");

    CURLcode result = curl_easy_perform(curl);
    if(result != CURLE_OK) {
        fprintf(stderr, "download problem: %s\n", curl_easy_strerror(result));
    }
    curl_easy_cleanup(curl);
    return EXIT_SUCCESS;
}
