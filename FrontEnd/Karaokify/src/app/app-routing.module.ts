import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { SearchScreenItIsComponent } from "./search-screen-it-is/search-screen-it-is.component";
import { HomeScreenComponent } from "./home-screen/home-screen.component";
import {LyricsPageComponent} from "./lyrics-page/lyrics-page.component";
import {SongsPageComponent} from "./songs-page/songs-page.component";

const routes: Routes = [
  { path: '', component:HomeScreenComponent},
  { path: 'screen-search', component:SearchScreenItIsComponent},
  { path: 'songs', component:SongsPageComponent},
  { path: 'lyrics', component:LyricsPageComponent}
];


@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]

})
export class AppRoutingModule { }
export const routingComponents = [SearchScreenItIsComponent, LyricsPageComponent, SongsPageComponent]