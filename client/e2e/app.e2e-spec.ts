import { FintJavadocsPage } from './app.po';

describe('fint-javadocs App', () => {
  let page: FintJavadocsPage;

  beforeEach(() => {
    page = new FintJavadocsPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
