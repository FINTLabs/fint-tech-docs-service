import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MdCardModule } from '@angular/material';

import { AngularFontAwesomeModule } from 'angular-font-awesome/angular-font-awesome';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { RepoCardComponent } from './repo-card/repo-card.component';
import { RepoListComponent } from './repo-list/repo-list.component';
import { RepositoriesService } from './api/repositories.service';


@NgModule({
  declarations: [
    AppComponent,
    RepoCardComponent,
    RepoListComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AngularFontAwesomeModule,
    MdCardModule,
    
    AppRoutingModule
  ],
  providers: [RepositoriesService],
  bootstrap: [AppComponent]
})
export class AppModule { }
