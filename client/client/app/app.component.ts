import { Component, OnInit, EventEmitter } from '@angular/core';
import { RepositoriesService } from './api/repositories.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  isLoading = this.repo.isLoading;
  private _searchInput: string;
  get searchInput() { return this._searchInput; }
  set searchInput(v) { this._searchInput = v; this.searchChanged.emit(v); }
  searchChanged = new EventEmitter<string>();

  constructor(private repo: RepositoriesService) { }

  ngOnInit() { }
}
