import DashboardLayout from '@/modules/dashboard/layouts/DashboardLayout.vue'

export const dashboardRoutes = [
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { layout: 'dashboard' },
    children: [
      {
        path: '',
        name: 'dashboard-overview',
        component: () => import('@/modules/dashboard/views/OverviewView.vue'),
      },
    ],
  },
]
