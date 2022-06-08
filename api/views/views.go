package views

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
	"yaroslavl-parkings/api"
	"yaroslavl-parkings/api/views/pages"
	"yaroslavl-parkings/data/order"
	"yaroslavl-parkings/data/rate"
	"yaroslavl-parkings/data/sessionstorer"
	"yaroslavl-parkings/data/user"
	"yaroslavl-parkings/pkg/qiwi"
)

// dependencies
type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*sessionstorer.Session, bool)
}

type PricingDatabaseInterface interface {
	GetActiveHoursDiscount() (rate.PeriodDiscount, error)
	GetSluggishHoursDiscount() (rate.PeriodDiscount, error)
	GetAdultRatePerHour() (rate.BaseRate, error)
	GetRetireeRatePerHour() (rate.BaseRate, error)
}

type OrderDatabaseInterface interface {
	GetOrderByID(id uint) (*order.Order, error)
	GetAllOrders() []order.Order
	GetAllOrdersByUserID(uid uint) []order.Order
	MarkOrderAsPaid(id uint) error
}

type viewsDependencies struct {
	pages      *pages.Pages
	sessions   SessionsInterface
	princingDB PricingDatabaseInterface
	ordersDB   OrderDatabaseInterface
	paymenter  *qiwi.Paymenter
}

type Views struct {
	SignIn                      http.HandlerFunc
	SignUp                      http.HandlerFunc
	MakeOrder                   http.HandlerFunc
	AddParkingPlace             http.HandlerFunc
	RemoveParkingPlace          http.HandlerFunc
	Profile                     http.HandlerFunc
	SeePricing                  http.HandlerFunc
	ChangeActiveHoursDiscount   http.HandlerFunc
	ChangeSluggishHoursDiscount http.HandlerFunc
	ChangeAdultBaseRate         http.HandlerFunc
	ChangeRetireeBaseRate       http.HandlerFunc
	PaymentPage                 http.HandlerFunc
	OrdersPage                  http.HandlerFunc
}

func NewViews(
	pathToTemplates string,
	sessions SessionsInterface,
	pricingDB PricingDatabaseInterface,
	orderDB OrderDatabaseInterface,
	paymenter *qiwi.Paymenter,
) *Views {
	dependencies := &viewsDependencies{
		pages:      pages.NewPages(pathToTemplates),
		sessions:   sessions,
		princingDB: pricingDB,
		ordersDB:   orderDB,
		paymenter:  paymenter,
	}

	return &Views{
		SignIn:                      dependencies.SignIn,
		SignUp:                      dependencies.SignUp,
		MakeOrder:                   dependencies.MakeOrder,
		AddParkingPlace:             dependencies.AddParkingPlace,
		RemoveParkingPlace:          dependencies.RemoveParkingPlace,
		Profile:                     dependencies.ProfilePage,
		SeePricing:                  dependencies.SeePricingPage,
		ChangeActiveHoursDiscount:   dependencies.ChangeActiveHoursDiscountPage,
		ChangeSluggishHoursDiscount: dependencies.ChangeSluggishHoursDiscountPage,
		ChangeAdultBaseRate:         dependencies.ChangeAdultBaseRatePage,
		ChangeRetireeBaseRate:       dependencies.ChangeRetireeBaseRatePage,
		PaymentPage:                 dependencies.PaymentPage,
		OrdersPage:                  dependencies.OrdersPage,
	}
}

// SignIn - serves SignIn page
func (d *viewsDependencies) SignIn(w http.ResponseWriter, r *http.Request) {
	if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		fmt.Println(d.pages.Public.SignIn.Execute(w, nil))
	}
}

// SignUp - serves SignUp page
func (d *viewsDependencies) SignUp(w http.ResponseWriter, r *http.Request) {
	if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		fmt.Println(d.pages.Public.SignUp.Execute(w, nil))
	}
}

// MakeOrder - serves MakeOrder page
func (d *viewsDependencies) MakeOrder(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.MakeOrder.Execute(w, nil))
	} else if api.IsAuth(d.sessions, r) {
		fmt.Println(d.pages.Private.MakeOrder.Execute(w, nil))
	} else {
		// redirection to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// ProfilePage - serves Profile page
func (d *viewsDependencies) ProfilePage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.Profile.Execute(w, nil))
	} else if api.IsAuth(d.sessions, r) {
		fmt.Println(d.pages.Private.Profile.Execute(w, nil))
	} else {
		// redirection to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// AddParkingPlace - serves AddParkingPlace page
func (d *viewsDependencies) AddParkingPlace(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.AddParking.Execute(w, nil))
	} else if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		// redirect to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// RemoveParkingPlace - serves RemoveParkingPlace page
func (d *viewsDependencies) RemoveParkingPlace(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.RemoveParking.Execute(w, nil))
	} else if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		// redirect to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// SeePricingPage - serves information about pricings
func (d *viewsDependencies) SeePricingPage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		activeHours, _ := d.princingDB.GetActiveHoursDiscount()
		sluggishHours, _ := d.princingDB.GetSluggishHoursDiscount()
		adultRate, _ := d.princingDB.GetAdultRatePerHour()
		retireeRate, _ := d.princingDB.GetRetireeRatePerHour()

		fmt.Println(d.pages.Admin.SeePricing.Execute(w,
			struct {
				ActiveHoursDiscount  uint
				SlugishHoursDiscount uint
				AdultHourlyPrice     uint
				RetireeHourlyRate    uint
			}{
				ActiveHoursDiscount:  uint(activeHours.DiscountInPercents),
				SlugishHoursDiscount: uint(sluggishHours.DiscountInPercents),
				AdultHourlyPrice:     uint(adultRate.Price),
				RetireeHourlyRate:    uint(retireeRate.Price),
			}))
		return
	}

	// redirect to the page for unauthenticated users
	http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}

// ChangeAdultBaseRate - changes adult base rate
func (d *viewsDependencies) ChangeAdultBaseRatePage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.ChangeAdultBaseRate.Execute(w, nil))
		return
	}

	// redirect to the page for unauthenticated users
	http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}

// ChangeRetireeBaseRate - changes retiree base rate
func (d *viewsDependencies) ChangeRetireeBaseRatePage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.ChangeRetireeBaseRate.Execute(w, nil))
		return
	}

	// redirect to the page for unauthenticated users
	http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}

// ChangeActiveHoursDiscount - changes active hours discount
func (d *viewsDependencies) ChangeActiveHoursDiscountPage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.ChangeActiveHoursDiscount.Execute(w, nil))
		return
	}

	// redirect to the page for unauthenticated users
	http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}

// ChangeSluggishHoursDiscountPage - changes sluggish hours discount
func (d *viewsDependencies) ChangeSluggishHoursDiscountPage(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.ChangeSluggishHoursDiscount.Execute(w, nil))
		return
	}

	// redirect to the page for unauthenticated users
	http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}

func (d *viewsDependencies) PaymentPage(w http.ResponseWriter, r *http.Request) {
	if !api.IsAuth(d.sessions, r) {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	orderID := r.URL.Query().Get("orderID")
	if orderID == "" {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	order, err := d.ordersDB.GetOrderByID(uint(orderIDInt))
	if err != nil {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	data := struct {
		Link string
		Bill uint
	}{
		Link: order.PaymentURL,
		Bill: order.Sum,
	}

	if api.IsAuthAndAdmin(d.sessions, r) {
		d.pages.Admin.PaymentPage.Execute(w, data)
	} else if api.IsAuth(d.sessions, r) {
		d.pages.Private.PaymentPage.Execute(w, data)
		fmt.Println(d.pages.Private.MakeOrder.Execute(w, nil))
	} else {
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}

}

func (d *viewsDependencies) OrdersPage(w http.ResponseWriter, r *http.Request) {
	session, valid := api.GetSessionIfValid(d.sessions, r)
	if !valid {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	var data struct {
		Orders   []order.Order
		Users    []user.User
		Username string
	}

	if api.IsAuthAndAdmin(d.sessions, r) {
		data.Orders = d.ordersDB.GetAllOrders()

		var wg sync.WaitGroup

		for _, o := range data.Orders {
			data.Users = append(data.Users, o.User)
			// request order status only if the payment time is not over
			if o.PaymentTimeout.After(time.Now()) {
				continue
			} else {
				wg.Add(1)
				go func(o order.Order, wg *sync.WaitGroup) {
					status, err := d.paymenter.GetBillStatus(o.StringID)
					if status == qiwi.PAID {
						d.ordersDB.MarkOrderAsPaid(o.ID)
					}
					if err == nil {
						o.Status = order.OrderStatus(status)
					}
					wg.Done()
				}(o, &wg)
			}
		}

		wg.Wait()

		d.pages.Admin.Orders.Execute(w, data)
	} else if api.IsAuth(d.sessions, r) {
		data.Orders = d.ordersDB.GetAllOrdersByUserID(session.UserID)
		data.Username = session.User.Username
		d.pages.Private.Orders.Execute(w, data)
	} else {
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}
