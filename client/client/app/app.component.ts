import { Component } from '@angular/core';

import { Http } from "@angular/http";
import { Observable, Observer, ReplaySubject } from 'rxjs/Rx';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private http: Http) { }

  isLoading = false;
  repositories = [];
  ngOnInit(){
    this.getRepositories();
    //console.log(this.repositories);
  }
  getRepositories() {
    console.log('getRepositories');
    /*
      if (evt) {
        evt.preventDefault();
        evt.stopPropagation();
      }
      */
      this.isLoading = true;
      this.http.get(`/api/projects`)
        .map(result => result.json())
        .catch(error => {console.error(error); return error; })
        .subscribe((result: any) => {
          this.repositories = result;
          console.log(result);
          this.isLoading = false;
        });
  }
}
