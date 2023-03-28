import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-lyrics-page',
  templateUrl: './lyrics-page.component.html',
  styleUrls: ['./lyrics-page.component.css']
})
export class LyricsPageComponent implements OnInit {

  youtubeUrl: string = "https://www.youtube.com/watch?v=S2Cti12XBw4";
  videoId: string | undefined;
  safeUrl: SafeResourceUrl | undefined;
  lyrics: string | undefined;
  loading: boolean = true;

  constructor(
    private http: HttpClient,
    private _sanitizer: DomSanitizer
  ) { }

  ngOnInit(): void {
    this.videoId = this.getVideoId(this.youtubeUrl);
    this.safeUrl = this._sanitizer.bypassSecurityTrustResourceUrl("https://www.youtube.com/embed/" + this.videoId);

    this.loadLyrics();
  }

  getVideoId(url: string) {
    let videoId = '';
    if (url.indexOf("v=") > -1) {
      videoId = url.split("v=")[1];
      let ampPos = videoId.indexOf("&");
      if (ampPos > -1) {
        videoId = videoId.substring(0, ampPos);
      }
    } else {
      videoId = url.split("/").pop() ?? '';
    }
    return videoId;
  }

  stripHtmlTags(html: string) {
    return html.replace(/<\/?[^>]+(>|$)/g, "");
  }

  loadLyrics() {
    let url = `https://video.google.com/timedtext?lang=en&v=${this.videoId}`;
    this.http.get(url, { responseType: 'text' }).subscribe((data) => {
      const parser = new DOMParser();
      const xmlDoc = parser.parseFromString(data, "text/xml");
      const textNodes = xmlDoc.getElementsByTagName('text');
      let lyrics = '';
      for (let i = 0; i < textNodes.length; i++) {
        lyrics += textNodes[i].textContent + ' ';
      }
      this.lyrics = this.stripHtmlTags(lyrics);
      this.loading = false;
    });
  }
}


