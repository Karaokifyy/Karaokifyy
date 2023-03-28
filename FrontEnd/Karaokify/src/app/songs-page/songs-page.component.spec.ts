import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { SongsPageComponent } from './songs-page.component';

describe('SongsPageComponent', () => {
  let component: SongsPageComponent;
  let fixture: ComponentFixture<SongsPageComponent>;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SongsPageComponent],
      providers: [
        { provide: Router, useValue: { navigate: jasmine.createSpy('navigate') } }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(SongsPageComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);
  });

  it('should navigate to the next page when the button is clicked', () => {
    // find the button element and trigger a click event
    const button = fixture.nativeElement.querySelector('button');
    button.click();

    // assert that the router.navigate method was called with the expected arguments
    (expect as any)(router.navigate).toHaveBeenCalledWith(['/next-page']);
  });
});